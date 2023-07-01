// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import './SEN.sol';
import 'hardhat/console.sol';

import './security/Verifier.sol';
import './security/VerifySignature.sol';

contract SentinelsBridge {
    // using ECDSA for bytes32;
    SEN private sentinels;
    constructor(address addr) {
        sentinels = SEN(addr);
        sentinels.setBridge(address(this));
    }

    struct Account {
        bool exsist;
        mapping(uint256 => bool) token_exsist;
    }
    mapping(address => Account) account;

    event Depositor(address deposit, uint256 tokenID, string tokenURL);
    event Withdrawer(address withdrawer, uint256 tokenID);

    function verify(
        string calldata message,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) public view returns (bool) {
        return verifyString(message, _v, _r, _s);
    }

    // function verifySignature(bytes32 _message, bytes memory signature) public view returns (bool) {
    //     console.log("msg.sender verifySignature: ", msg.sender);
    //     console.log("recover verifySignature: ", _message.recover(signature));
    //     return _message.recover(signature) == msg.sender;
    // }

    function deposit(uint256 tokenId) public {
        require(sentinels.ownerOf(tokenId) == msg.sender, "only owner of token can deposit");
        string memory tokenURI = sentinels.tokenURI(tokenId);
        account[msg.sender].exsist = true;
        account[msg.sender].token_exsist[tokenId] = true;
        lock(msg.sender, tokenId);
        emit Depositor(msg.sender, tokenId, tokenURI);
    }

    function withdraw(uint256 tokenID, string memory url) public {
        if (account[msg.sender].exsist && account[msg.sender].token_exsist[tokenID]){
            unlock(msg.sender, tokenID);
            account[msg.sender].exsist = false;
            account[msg.sender].token_exsist[tokenID] = false;
            emit Withdrawer(msg.sender, tokenID);
            return;
        }
        account[msg.sender].exsist = true;
        account[msg.sender].token_exsist[tokenID] = true;
        sentinels.mintProtocol(msg.sender, tokenID, url);
        emit Withdrawer(msg.sender, tokenID);
        return;
    }
    
    function lock(address _to, uint256 tokenId) private {
        sentinels.transferItemTo(_to, address(this), tokenId);
    }

    function unlock(address _to, uint256 tokenId) private {
        sentinels.transferItemTo(address(this), _to, tokenId);
    }

    function verifyString(string memory message, uint8 v, bytes32 r, bytes32 s) 
        private view returns (bool) {

        // The message header; we will fill in the length next
        string memory header = "\x19Ethereum Signed Message:\n000000";

        uint256 lengthOffset;
        uint256 length;
        assembly {
            // The first word of a string is its length
            length := mload(message)
            // The beginning of the base-10 message length in the prefix
            lengthOffset := add(header, 57)
        }

        // Maximum length we support
        require(length <= 999999);

        // The length of the message's length in base-10
        uint256 lengthLength = 0;

        // The divisor to get the next left-most message length digit
        uint256 divisor = 100000;

        // Move one digit of the message length to the right at a time
        while (divisor != 0) {

            // The place value at the divisor
            uint256 digit = length / divisor;
            if (digit == 0) {
                // Skip leading zeros
                if (lengthLength == 0) {
                    divisor /= 10;
                    continue;
                }
            }

            // Found a non-zero digit or non-leading zero digit
            lengthLength++;

            // Remove this digit from the message length's current value
            length -= digit * divisor;

            // Shift our base-10 divisor over
            divisor /= 10;

            // Convert the digit to its ASCII representation (man ascii)
            digit += 0x30;
            // Move to the next character and write the digit
            lengthOffset++;

            assembly {
                mstore8(lengthOffset, digit)
            }
        }

        // The null string requires exactly 1 zero (unskip 1 leading 0)
        if (lengthLength == 0) {
            lengthLength = 1 + 0x19 + 1;
        } else {
            lengthLength += 1 + 0x19;
        }

        // Truncate the tailing zeros from the header
        assembly {
            mstore(header, lengthLength)
        }

        // Perform the elliptic curve recover operation
        bytes32 check = keccak256(abi.encodePacked(header, message));

        return ecrecover(check, v, r, s) == msg.sender;
    }
}