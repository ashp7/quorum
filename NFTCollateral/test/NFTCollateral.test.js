const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("NFTCollateralLoan", function () {
    let NFTCollateralLoan, nftCollateralLoan, MockNFT, mockNFT;
    let owner, addr1;

    beforeEach(async function () {

        [owner, addr1] = await ethers.getSigners();

        MockNFT = await ethers.getContractFactory("MockNFT");
        NFTCollateralLoan = await ethers.getContractFactory("NFTCollateralLoan");

        tokenId = 100;
        mockNFT = await MockNFT.deploy();
        await mockNFT.connect(owner).mint(addr1.address, tokenId);
        await mockNFT.mint(addr1.address,BigInt(1)); // Mint an NFT to addr1
        console.log("addr1 address:", addr1)

        nftCollateralLoan = await NFTCollateralLoan.deploy();

        console.log("MockNFT address:", mockNFT.target);
        console.log("NFTCollateralLoan address:", nftCollateralLoan.target);

    });

    describe("submitProposal", function () {
        it("should allow a user to submit a loan proposal", async function () {
            // Transfer the NFT to the contract from addr1
            console.log("first connect");
          await mockNFT.connect(addr1.target).approve(nftCollateralLoan.target, tokenId);

            console.log("submit one proposal");
            // `addr1` submits a proposal
            loanAmount = 10000;
            interestRate = 10
            duration = 36000
            await nftCollateralLoan.connect(addr1.target).submitProposal(mockNFT.target, tokenId, loanAmount, interestRate, duration);

            // Retrieve the first proposal
            const proposalIndex = 0; // Assuming this is the first proposal
            const proposal = await nftCollateralLoan.proposals(proposalIndex);

            expect(proposal.borrower).to.equal(addr1.address);
            expect(proposal.nftContractAddress).to.equal(mockNFT.address);
            expect(proposal.nftTokenId).to.equal(tokenId);
            expect(proposal.loanAmount).to.equal(loanAmount);
            // Additional assertions as necessary
        });
    });

    // Additional tests can be added here
});

