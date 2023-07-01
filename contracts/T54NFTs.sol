// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "hardhat/console.sol";

contract T54NFTs is ERC721URIStorage {
    mapping(uint256 => bool) private tokenMap;
    mapping(uint256 => bool) private synctoken;
    uint256[] private tid;
    mapping(uint256 => address) private fromT55Deposit;

    constructor() ERC721("T54", "T54") {}

    event MintNFTs(address player, uint256 tokenID, string url);

    function mint(address player, string memory tokenURI) public returns(uint256) {
        uint256 tokenID = 0;
        while (true) {
            if (tokenMap[tokenID] == false && synctoken[tokenID] == false) {
                tid.push(tokenID);
                tokenMap[tokenID] = true;
                break;
            }
            tokenID = tokenID + 1;
        }
        uint256 newItemId = tid[tid.length - 1];
        _mint(player, newItemId);
        _setTokenURI(newItemId, tokenURI);
        emit MintNFTs(player, tokenID, tokenURI);
        return newItemId;
    }

    function mintTo(address _to, uint256 tokenID, string memory url) public {
        require(tokenMap[tokenID] == false, "Requires TokenID unique");
        require(synctoken[tokenID] == true, "Require synctoken is true");
        fromT55Deposit[tokenID] = address(0);
        tokenMap[tokenID] == true;
        _mint(_to, tokenID);
        _setTokenURI(tokenID, url);
        tid.push(tokenID);
    }

    function getTokenID() public view returns (uint256) {
        return tid[tid.length - 1];
    }

    function sync(uint256 tokenId) public {
        synctoken[tokenId] = true;
    }

    function transfer(address _from, address _to, uint256 _tokenID) public {
        require(ownerOf(_tokenID) == _from, "You don't own this NFTs");
        _approve(msg.sender, _tokenID);
        transferFrom(_from, _to, _tokenID);
    }
}
