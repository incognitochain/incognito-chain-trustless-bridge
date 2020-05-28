pragma solidity 0.6.6;
pragma experimental ABIEncoderV2;
contract BTCProof {

    struct MerkleProof {
        bytes32 proofHash;
        bool isLeft;
    }

    /**
     * @dev hash the merkle root
     */
    function hash256(bytes memory _b) internal pure returns (bytes32) {
        return sha256(abi.encodePacked(sha256(_b)));
    }

    /**
     * @dev hash the merkle node
     */
    function _hash256MerkleStep(bytes32 _a, bytes32 _b) internal pure returns (bytes32) {
        return hash256(abi.encodePacked(_a, _b));
    }

    /**
     * @dev verify bitcoin transaction
     * @param merkleRoot: the merkleRoot in blockchain
     * @param merkleProofs: nodes of merkle tree
     * @param txHash: the transaction ID need to verified
     * @return is transaction valid or not
     */
    function verify(
        bytes32 merkleRoot,
        MerkleProof[] memory merkleProofs,
        bytes32 txHash
    ) public pure returns (bool) {
        bytes32 curHash = txHash;
        for (uint i = 0; i < merkleProofs.length; i++) {
            if (merkleProofs[i].isLeft) {
                curHash = _hash256MerkleStep(merkleProofs[i].proofHash, curHash);
            } else {
                curHash = _hash256MerkleStep(curHash, merkleProofs[i].proofHash);
            }
        }
        return curHash == merkleRoot;
    }
}