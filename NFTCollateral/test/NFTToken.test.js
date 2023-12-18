const { expect } = require("chai");

describe("NFTToken", function () {
    let nftToken;
    let owner;
    let addr1;
    let addr2;

    beforeEach(async function () {
        const NFTToken = await ethers.getContractFactory("NFTToken");
        [owner, addr1, addr2] = await ethers.getSigners();

        nftToken = await NFTToken.deploy();
        await nftToken.waitForDeployment()
    });

    it("should mint correctly", async function () {
        const tokenURI = "";
        await nftToken.mint(addr1.address, tokenURI);

        const tokenId = 1;
        const [ownerAddress, retrievedURI] = await nftToken.getToken(tokenId);

        expect(ownerAddress).to.equal(addr1.address);
        expect(retrievedURI).to.equal(tokenURI);
    });
});
