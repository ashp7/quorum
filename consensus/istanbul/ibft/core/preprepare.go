// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"time"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	istanbulcommon "github.com/ethereum/go-ethereum/consensus/istanbul/common"
	ibfttypes "github.com/ethereum/go-ethereum/consensus/istanbul/ibft/types"
)

func (c *core) sendPreprepare(request *istanbul.Request) {
	logger := c.logger.New("state", c.state, "address", c.Address(), "round", c.current.Round().String(), "request proposal", request.Proposal.Number(), "method", "sendPreprepare")


	// If I'm the proposer and I have the same sequence with the proposal
	if c.current.Sequence().Cmp(request.Proposal.Number()) == 0 && c.IsProposer() {
		logger.Info("I am the proposer and I have the same sequence number as the request")
		curView := c.currentView()
		preprepare, err := ibfttypes.Encode(&istanbul.Preprepare{
			View:     curView,
			Proposal: request.Proposal,
		})
		if err != nil {
			logger.Error("Failed to encode", "view", curView)
			return
		}
		c.broadcast(&ibfttypes.Message{
			Code: ibfttypes.MsgPreprepare,
			Msg:  preprepare,
		})
		logger.Info("Broadcasting preprepare message")
	}
}

func (c *core) handlePreprepare(msg *ibfttypes.Message, src istanbul.Validator) error {
	logger := c.logger.New("from", src,"state", c.state, "address", c.Address(), "round", c.current.Round().String(), "message", msg.String())


	logger.Info("Decoding prepare message")
	// Decode PRE-PREPARE
	var preprepare *istanbul.Preprepare
	err := msg.Decode(&preprepare)
	if err != nil {
		return istanbulcommon.ErrFailedDecodePreprepare
	}

	logger.Info("Checking the current view with the message view")
	// Ensure we have the same view with the PRE-PREPARE message
	// If it is old message, see if we need to broadcast COMMIT
	if err := c.checkMessage(ibfttypes.MsgPreprepare, preprepare.View); err != nil {
		if err == istanbulcommon.ErrOldMessage {
			// Get validator set for the given proposal
			valSet := c.backend.ParentValidators(preprepare.Proposal).Copy()
			previousProposer := c.backend.GetProposer(preprepare.Proposal.Number().Uint64() - 1)
			valSet.CalcProposer(previousProposer, preprepare.View.Round.Uint64())
			// Broadcast COMMIT if it is an existing block
			// 1. The proposer needs to be a proposer matches the given (Sequence + Round)
			// 2. The given block must exist
			if valSet.IsProposer(src.Address()) && c.backend.HasPropsal(preprepare.Proposal.Hash(), preprepare.Proposal.Number()) {
				c.sendCommitForOldBlock(preprepare.View, preprepare.Proposal.Hash())
				return nil
			}
		}
		return err
	}

	logger.Info("Checking if message comes from current proposer")
	// Check if the message comes from current proposer
	if !c.valSet.IsProposer(src.Address()) {
		logger.Warn("Ignore preprepare messages from non-proposer")
		return istanbulcommon.ErrNotFromProposer
	}



	logger.Info("Checking the proposal", "proposal number", preprepare.Proposal.Number())
	// Verify the proposal we received
	if duration, err := c.backend.Verify(preprepare.Proposal); err != nil {
		// if it's a future block, we will handle it again after the duration
		if err == consensus.ErrFutureBlock {
			logger.Info("Proposed block will be handled in the future", "err", err, "duration", duration)
			c.stopFuturePreprepareTimer()
			c.futurePreprepareTimer = time.AfterFunc(duration, func() {
				c.sendEvent(backlogEvent{
					src: src,
					msg: msg,
				})
			})
		} else {
			logger.Warn("Failed to verify proposal", "err", err, "duration", duration)
			c.sendNextRoundChange()
		}
		return err
	}

	// Here is about to accept the PRE-PREPARE
	if c.state == ibfttypes.StateAcceptRequest {
		// Send ROUND CHANGE if the locked proposal and the received proposal are different
		if c.current.IsHashLocked() {
			if preprepare.Proposal.Hash() == c.current.GetLockedHash() {
				// Broadcast COMMIT and enters Prepared state directly
				logger.Info("Locked hash Accepting pre-prepare message, setting state to prepared and sending commit directly")
				c.acceptPreprepare(preprepare)
				c.setState(ibfttypes.StatePrepared)
				c.sendCommit()
			} else {
				// Send round change
				c.sendNextRoundChange()
			}
		} else {
			// Either
			//   1. the locked proposal and the received proposal match
			//   2. we have no locked proposal
			logger.Info("Accepting pre-prepare message, setting state to preprepared and sending prepare")
			c.acceptPreprepare(preprepare)
			c.setState(ibfttypes.StatePreprepared)
			c.sendPrepare()
		}
	}

	return nil
}

func (c *core) acceptPreprepare(preprepare *istanbul.Preprepare) {
	c.consensusTimestamp = time.Now()
	c.current.SetPreprepare(preprepare)
}
