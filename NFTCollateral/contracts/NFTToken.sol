// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";          // Base ERC721 contract
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";  // ERC721 extension for token URI storage

contract NFTToken is ERC721URIStorage {
    uint256 private _nextTokenId;

 // Event declaration
    event NFTMinted(uint256 indexed tokenId, address indexed recipient, string tokenURI);

    constructor() ERC721("NFTToken", "NFTN") {
        _nextTokenId = 1; // Initialize the next token ID
    }

    function mint(address recipient, string memory tokenURI) public returns (uint256) {
        uint256 tokenId = _nextTokenId++;
        _mint(recipient, tokenId);
        _setTokenURI(tokenId, tokenURI);

     // Emit the event
        emit NFTMinted(tokenId, recipient, tokenURI);

        return tokenId;
    }

    function getToken(uint256 tokenId) public view returns (address owner, string memory uri) {
	owner = ownerOf(tokenId);
	uri = tokenURI(tokenId);
    }
}
