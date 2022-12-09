// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

contract Album {
    struct AlbumInfo {
        int256 price;
        string creater;
        string title;
    }
    AlbumInfo[] public albumInfo;
    address public addressAfterVerify;
    address public messageToAddress;

    function setInfo(
        int256 _price,
        string memory _title,
        string memory _creater,
        string memory message,
        string memory sig
    ) public {
        require(verify(message, sig),"Function must call from backend");
        albumInfo.push(AlbumInfo(_price, _title, _creater));
    }

    function getInfo(uint256 index) public view returns (AlbumInfo memory) {
        return albumInfo[index];
    }

    function verify(string memory message, string memory sig) public returns (bool) {
        // bytes32 message = ethMessageHash("TEST");

        // bytes
        //     memory sig = hex"bceab59162da5e511fb9c37fda207d443d05e438e5c843c57b2d5628580ce9216ffa0335834d8bb63d86fb42a8dd4d18f41bc3a301546e2c47aa1041c3a1823701";
        address addr = 0x9105C6A78Cf3CA31e1F2615665efBfc01263532B;
        return recover(stringToBytes32(message),string_tobytes(sig)) == addr;
    }

    /**
     * @dev Recover signer address from a message by using their signature
     * @param hash bytes32 message, the hash is the signed message. What is recovered is the signer address.
     * @param sig bytes signature, the signature is generated using web3.eth.sign()
     */
    function recover(
        bytes32 hash,
        bytes memory sig
    ) internal returns (address) {
        bytes32 r;
        bytes32 s;
        uint8 v;

        // Check the signature length
        // if (sig.length != 65) {
        //     return (address(0));
        // }

        // Divide the signature in r, s and v variables
        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
        recoverSignerFromSignature(v, r, s,hash);

        // Version of signature should be 27 or 28, but 0 and 1 are also possible versions
        if (v < 27) {
            v += 27;
        }

        // If the version is correct return the signer address
        if (v != 27 && v != 28) {
            return (address(0));
        } else {
            // solium-disable-next-line arg-overflow
            addressAfterVerify=ecrecover(hash, v, r, s);
            return ecrecover(hash, v, r, s);
        }
    }


    /**
     * @dev prefix a bytes32 value with "\x19Ethereum Signed Message:" and hash the result
     */
    function ethMessageHash(
        string memory message
    ) internal pure returns (bytes32) {
        return
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(abi.encodePacked(message))));
    }

    function stringToBytes32(string memory source) public pure returns (bytes32 result) {
    bytes memory tempEmptyStringTest = bytes(source);
    if (tempEmptyStringTest.length == 0) {
        return 0x0;
    }

    assembly {
        result := mload(add(source, 32))
    }
}

 function string_tobytes( string memory s)public returns (bytes memory){
        return bytes(s);
    }


function recoverSignerFromSignature(uint8 v, bytes32 r, bytes32 s, bytes32 hash) public {
    address signer = ecrecover(hash, v, r, s);
    require(signer != address(0), "ECDSA: invalid signature");
}

}
