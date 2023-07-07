// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import './LEO.sol';
import './security/Verifier.sol';

contract LeopardBridge is Verifier {
    using ECDSA for bytes32;
    using ECDSA for bytes;
    LEO private leopard;
    constructor(address addr) {
        leopard = LEO(addr);
        leopard.setBridge(address(this));
    }

    struct Account {
        bool exsist;
        mapping(uint256 => bool) token_exsist;
    }
    mapping(address => Account) account;

    event Depositor(address deposit, uint256 tokenID, string tokenURL);
    event Withdrawer(address withdrawer, uint256 tokenID);

    function SignVerify(string memory message, bytes memory signature) public view returns (bool) {
        return bytes(message).toEthSignedMessageHash().recover(signature) == msg.sender;
    }

    function verify(bytes memory signature, address owner, uint256 tokenId) private pure returns(bool) {
        bytes32 message = keccak256(abi.encodePacked(owner, tokenId));
        bytes32 hash = message.toEthSignedMessageHash();
        return hash.recover(signature) == owner;
    }

    function deposit(uint256 tokenId) public {
        require(leopard.ownerOf(tokenId) == msg.sender, "only owner of token can deposit");
        string memory tokenURI = leopard.tokenURI(tokenId);
        account[msg.sender].exsist = true;
        account[msg.sender].token_exsist[tokenId] = true;
        lock(msg.sender, tokenId);
        emit Depositor(msg.sender, tokenId, tokenURI);
    }

    function withdraw(uint256 tokenID) public {
        console.log("withdraw: ",msg.sender);
        require(account[msg.sender].exsist == true, "only owner of token can withdraw");
        require(account[msg.sender].token_exsist[tokenID] == true, "tokenID must be existing");
        unlock(msg.sender, tokenID);
        account[msg.sender].exsist = false;
        account[msg.sender].token_exsist[tokenID] = false;
        emit Withdrawer(msg.sender, tokenID);
    }

    function lock(address _to, uint256 tokenId) private {
        leopard.transferItemTo(_to, address(this), tokenId);
    }

    function unlock(address _to, uint256 tokenId) private {
        leopard.transferItemTo(address(this), _to, tokenId);
    }
}