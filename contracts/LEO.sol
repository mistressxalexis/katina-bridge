// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "hardhat/console.sol";

contract LEO is ERC721URIStorage {
    using ECDSA for bytes32;
    address private admin;
    address private bridge;
    constructor() ERC721("Leopard", "LEO") {
        admin = msg.sender;
        id = 0;
        bridge = address(0);
    }
    struct Instance {
        bool exists;
        bool isOwner;
        address addr;
    }
    uint256 [] private NFTID;
    uint256 private id;
    mapping(uint256 => Instance) private instance;

    function setBridge(address _bridge) public {
        if (bridge == address(0)) {
            bridge = _bridge;
        }
    }

    function signature(uint256 tokenId, string memory url) public view returns(bytes32) {
        bytes32 message = keccak256(abi.encodePacked(msg.sender, tokenId, url));
        bytes32 hash = message.toEthSignedMessageHash();
        return hash;
    }

    function mintItem(string memory url) public {
        uint256 tokenID = 0;
        while (true) {
            if (!instance[tokenID].exists) {
                id = tokenID;
                NFTID.push(tokenID);
                break;
            }
            tokenID = tokenID + 1;
        }
        _mint(msg.sender, tokenID);
        _setTokenURI(tokenID, url);
        instance[tokenID].exists = true;
        instance[tokenID].isOwner = true;
    }

    function mintProtocol(address _to, uint256 tokenId, string memory url) public {
        // console.log("bridge: " , bridge);
        // console.log("msg.sender: " , msg.sender);
        // console.log("_to: " , _to);
        require(_to == bridge, "requires bridge address");
        require(bridge == msg.sender, "only bridge can mint nfts");
        id = tokenId;
        NFTID.push(tokenId);
        _mint(_to, tokenId);
        _setTokenURI(tokenId, url);
        instance[tokenId].exists = true;
        instance[tokenId].isOwner = true;
    }

    function getTokenID() public view returns(uint256) {
        require(NFTID.length > 0, "nft is not available in address");
        return id;
    }

    function transferItemTo(address _from, address _to, uint256 tokenID) public {
        require(ownerOf(tokenID) == _from, "requried owner of token");
        _transfer(_from, _to, tokenID);
    }

    function tokenExists(uint256 tokenId) public view returns(bool){
        return _exists(tokenId);
    }

}
