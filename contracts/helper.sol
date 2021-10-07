pragma solidity ^0.8.0;

library helper {
    function verify(bytes32 hash, uint8 v, bytes32 r, bytes32 s, address signer)private pure returns(bool) {
    	bytes memory prefix = "\x19Ethereum Signed Message:\n32";
    	bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, hash));
    	return ecrecover(prefixedHash, v, r, s) == signer;
	}

    function verifySig(bytes32 hash, bytes[] calldata sig, address signer) public pure returns (bool) {
        return true;
    }
}