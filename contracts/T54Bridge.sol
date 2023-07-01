// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./T54NFTs.sol";

contract T54Bridge {
    struct Account {
        uint256 t54;
        bool exists;
    }
    mapping(address => Account) public accounts;
    mapping(uint256 => string) nfts;

    T54NFTs public t54;

    constructor(address _nfts) {
        t54 = T54NFTs(_nfts);
    }

    event Withdrawer(address from, address to, uint256 tokenID);
    event Depositor(address from, address to, uint256 tokenID, string url);
    event ClaimNFTs(address _to, uint256 tokenID, string url);

    function deposit(uint256 tokenID) external payable {
        accounts[msg.sender].t54 = tokenID;
        accounts[msg.sender].exists = true;
        lockNFT(tokenID);
        string memory url = t54.tokenURI(tokenID);
        emit Depositor(msg.sender, address(this), tokenID, url);
    }

    function withdraw(uint256 tokendID) external {
        require(accounts[msg.sender].exists = true, "depositor must exist");
        unlockNFT(tokendID);
        accounts[msg.sender].exists = false;
        accounts[msg.sender].t54 = type(uint256).max;
        emit Withdrawer(address(this), msg.sender, tokendID);
    }

    function claim(address _to, uint256 _tokenID, string memory url) external {
        t54.mintTo(_to, _tokenID, url);
    }
    function lockNFT(uint256 tokenID) private {
        require(
            t54.ownerOf(tokenID) == msg.sender,
            "Only the onwer can lock NFT"
        );
        t54.transfer(msg.sender, address(this), tokenID);
    }

    function unlockNFT(uint256 _tokenID) private {
        require(
            accounts[msg.sender].t54 == _tokenID,
            "Only the onwer can unlock NFT"
        );
        t54.transfer(address(this), msg.sender, _tokenID);
    }

    function getBridgeAddr() public view returns(address){
        return address(this);
    }

    function claimNFTs(address sender, uint256 tokenID) public {
        require(accounts[sender].t54 == tokenID, "Only the onwer can claim NFT");
        require(accounts[sender].exists == true, "Owner must be exists");
        emit ClaimNFTs(sender, tokenID, nfts[tokenID]);
    }

    function syncNFTs(address sender, uint256 tokenID, string memory url) public {
        accounts[sender].t54 = tokenID;
        accounts[sender].exists = true;
        nfts[tokenID] = url;
    }
}
