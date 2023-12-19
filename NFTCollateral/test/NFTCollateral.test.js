const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("NFTCollateralLoan", function () {
    let nftCollateral;
    let nftToken;
    let borrower;
    let lender;
    let col_accounts;
    let nft_owner;
    let col_owner;
    let nftContractAddress;
    let nftCollateralAddress;

    beforeEach(async function () {

        const NFTToken = await ethers.getContractFactory("NFTToken");
        const NFTCollateral = await ethers.getContractFactory("NFTCollateralLoan");

        [nft_owner, nft_addr1, nft_addr2, ...nft_accounts] = await ethers.getSigners();
        [col_owner, col_addr1, col_addr2, ...col_accounts] = await ethers.getSigners();

        nftToken = await NFTToken.deploy();
        await nftToken.waitForDeployment();

        nftCollateral = await NFTCollateral.deploy();
        await nftCollateral.waitForDeployment();

        nftContractAddress = await nftToken.getAddress()
        nftCollateralAddress = await nftCollateral.getAddress()

        // Set up a borrower and mint an NFT to them
        borrower = col_accounts[0];
        await nftToken.connect(nft_owner).mint(borrower.address, 1);


        lender = col_accounts[1];
        await nftToken.connect(nft_owner).mint(lender.address, 1);

        console.log("deployed addresses" , "NFT", nftContractAddress, "col", nftCollateralAddress)

    });

    it("should submit a proposal", async function () {
        const nftTokenId = 1;
        const loanAmount = 100; // Assuming loan amount is in ether
        const interestRate = 10;
        const duration = 30 * 24 * 60 * 60; // 30 days in seconds

        // Approve the NFTCollateral contract to transfer the NFT on behalf of the borrower
        await nftToken.connect(borrower).approve(nftCollateralAddress, nftTokenId);

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

    it("should accept a proposal", async function () {
        // Assuming a proposal has already been submitted
        const proposalId = 0;

        // Lender accepts the proposal
        await expect(nftCollateral.connect(lender).acceptProposal(proposalId))
            .to.emit(nftCollateral, "ProposalAccepted");

        // Retrieve the proposal and verify its properties
        const proposal = await nftCollateral.proposals(proposalId);
        expect(proposal.lender).to.equal(lender.address);
        expect(proposal.isAccepted).to.be.true;
    });


    it("should repay a loan", async function () {
        const proposalId = 0;
        const nftTokenId = 1;
        const loanAmount = 100; // Assuming loan amount is in ether
        const interestRate = 10;
        const duration = 30 * 24 * 60 * 60; // 30 days in seconds
        const loanAmountWei = ethers.parseUnits(loanAmount.toString(), "ether");
        const repaymentAmount = loanAmount * 2
        const repaymentAmountWei =  ethers.parseUnits(repaymentAmount.toString(), "ether");

        // Approve the NFTCollateral contract to transfer the NFT on behalf of the borrower
        await nftToken.connect(borrower).approve(nftCollateralAddress, nftTokenId);

        // Submit the proposal
        await expect(nftCollateral.connect(borrower).submitProposal(
            nftContractAddress,
            nftTokenId,
            loanAmountWei,
            interestRate,
            duration,
        )).to.emit(nftCollateral, "ProposalSubmitted"); // Check for ProposalSubmitted event

        // Accept the proposal
        await expect(nftCollateral.connect(lender).acceptProposal(proposalId, { value: loanAmountWei }))
            .to.emit(nftCollateral, "ProposalAccepted");


        const proposal1 = await nftCollateral.proposals(0);
        console.log("loan amount is", proposal1.loanAmount)


        // Set the borrower's balance to a high enough value
        const borrowerBalanceWei = ethers.parseUnits("1000", "ether"); // 1000 Ether in Wei
        await ethers.provider.send("hardhat_setBalance", [
            borrower.address,
            '0x' + borrowerBalanceWei.toString(16)
        ]);

        // Verify the borrower's balance (optional)
        const borrowerBalance = await ethers.provider.getBalance(borrower.address);
        console.log("Borrower's new balance:", borrowerBalance);

        await expect(nftCollateral.connect(borrower).repayLoan(proposalId, { value: repaymentAmountWei }))
            .to.emit(nftCollateral, "LoanRepaid")
            .withArgs(proposalId);

        // Verify the loan repayment status
        const proposal = await nftCollateral.proposals(proposalId);
        expect(proposal.isPaidBack).to.be.true;
    });

    it("should allow a borrower to retract a proposal", async function () {
        const nftTokenId = 1;
        const loanAmount = ethers.parseUnits("100", "ether");
        const interestRate = 10; // 10%
        const duration = 30 * 24 * 60 * 60; // 30 days

        // Submit a proposal
        await nftToken.connect(borrower).approve(nftCollateralAddress, nftTokenId);
        await nftCollateral.connect(borrower).submitProposal(
            nftContractAddress,
            nftTokenId,
            loanAmount,
            interestRate,
            duration
        );

        // Retract the proposal
        await expect(nftCollateral.connect(borrower).retractProposal(0))
            .to.emit(nftCollateral, "ProposalRetracted")
            .withArgs(0);
    });

    it("should transfer NFT to lender after loan date expiry", async function () {
        const proposalId = 0;
        const nftTokenId = 1;
        const loanAmount = ethers.parseUnits("100", "ether");
        const interestRate = 10; // 10%
        const duration = 30 * 24 * 60 * 60; // 30 days

        // Submit a proposal
        await nftToken.connect(borrower).approve(nftCollateralAddress, nftTokenId);
        await nftCollateral.connect(borrower).submitProposal(
            nftContractAddress,
            nftTokenId,
            loanAmount,
            interestRate,
            duration
        );

        // Accept the proposal
        await nftCollateral.connect(lender).acceptProposal(proposalId, { value: loanAmount });

        // Increase time
        await ethers.provider.send("evm_increaseTime", [duration + 1]);
        await ethers.provider.send("evm_mine", []);

        // Act on loan date expiry
        await nftCollateral.connect(col_owner).actOnLoanDateExpiry(proposalId);

        // Verify NFT ownership
        expect(await nftToken.ownerOf(nftTokenId)).to.equal(lender.address);
    });

    it("should retrieve a proposal", async function () {
        const nftTokenId = 1;
        const loanAmount = ethers.parseUnits("100", "ether");
        const interestRate = 10; // 10%
        const duration = 30 * 24 * 60 * 60; // 30 days

        // Submit a proposal
        await nftToken.connect(borrower).approve(nftCollateralAddress, nftTokenId);
        await nftCollateral.connect(borrower).submitProposal(
            nftContractAddress,
            nftTokenId,
            loanAmount,
            interestRate,
            duration
        );

        // Retrieve the proposal
        const proposal = await nftCollateral.getProposal(0);

        // Verify the proposal details
        expect(proposal.borrower).to.equal(borrower.address);
        expect(proposal.loanAmount).to.equal(loanAmount);
        // Add other necessary verifications
    });

    it("should return the total number of proposals", async function () {
        const nftTokenId = 1;
        const loanAmount = ethers.parseUnits("100", "ether");
        const interestRate = 10; // 10%
        const duration = 30 * 24 * 60 * 60; // 30 days

        // Submit a proposal
        await nftToken.connect(borrower).approve(nftCollateralAddress, nftTokenId);
        await nftCollateral.connect(borrower).submitProposal(
            nftContractAddress,
            nftTokenId,
            loanAmount,
            interestRate,
            duration
        );

        // Get the total number of proposals
        const totalProposals = await nftCollateral.getTotalProposals();

        // Verify the total number of proposals
        expect(totalProposals).to.equal(1);
    });

});
