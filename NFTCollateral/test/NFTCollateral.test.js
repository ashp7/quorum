const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("NFTCollateralLoan", function () {
    let nftCollateral;
    let nftToken;
    let borrower;
    let col_accounts;
    let nft_owner;
    let col_owner;

    beforeEach(async function () {
        // Deploy NFTToken
        const NFTToken = await ethers.getContractFactory("NFTToken");
        [nft_owner, nft_addr1, nft_addr2, ...nft_accounts] = await ethers.getSigners();
        nftToken = await NFTToken.deploy();
        await nftToken.waitForDeployment();

        // Deploy NFTCollateralLoan
        const NFTCollateral = await ethers.getContractFactory("NFTCollateralLoan");
        [col_owner, col_addr1, col_addr2, ...col_accounts] = await ethers.getSigners();
        nftCollateral = await NFTCollateral.deploy();
        await nftCollateral.waitForDeployment();

        // Set up a borrower and mint an NFT to them
        borrower = col_accounts[0];
        await nftToken.connect(nft_owner).mint(borrower.address, 1);
    });

    it("should submit a proposal", async function () {
        const nftContractAddress = nftToken.address; // Use the contract address of the NFTToken
        const nftTokenId = 1;
        const loanAmount = 100; // Assuming loan amount is in ether
        const interestRate = 10;
        const duration = 30 * 24 * 60 * 60; // 30 days in seconds

        // Approve the NFTCollateral contract to transfer the NFT on behalf of the borrower
        await nftToken.connect(borrower).approve(nftCollateral.address, nftTokenId);

        // Submit the proposal
        await expect(nftCollateral.connect(borrower).submitProposal(
            nftContractAddress,
            nftTokenId,
            loanAmount,
            interestRate,
            duration,
        )).to.emit(nftCollateral, "ProposalSubmitted"); // Check for ProposalSubmitted event

        // Retrieve the proposal and verify its properties
        const proposal = await nftCollateral.proposals(0);
        expect(proposal.borrower).to.equal(borrower.address);
        expect(proposal.nftContractAddress).to.equal(nftContractAddress);
        expect(proposal.nftTokenId).to.equal(nftTokenId);
        expect(proposal.loanAmount).to.equal(loanAmount);
        expect(proposal.interestRate).to.equal(interestRate);
        expect(proposal.duration).to.equal(duration);
        expect(proposal.isAccepted).to.be.false;
        expect(proposal.isPaidBack).to.be.false;
    });
});
