package main

import (
	"./btc"
	"./btc/btcrelay"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

var genesisAcc *account
var auth *bind.TransactOpts

func TestBuildBTCMerkleProof(t *testing.T) {
	bc :=  btcrelaying.GetBlockCypherAPI("test3")
	txID := "4478038c54fe1ea19668afc5f088861152cb35559f90ee39024b393a21a612cb"
	msgTx := btcrelaying.BuildMsgTxFromCypher(txID, "test3")

	cypherBlock, err := bc.GetBlock(
		1696574,
		"",
		map[string]string{
			"txstart": "0",
			"limit":   "500",
		},
	)

	require.Equal(t, nil, err)

	txIDs := cypherBlock.TXids
	txHashes := make([]*chainhash.Hash, len(txIDs))
	for i := 0; i < len(txIDs); i++ {
		txHashes[i], _ = chainhash.NewHashFromStr(txIDs[i])
	}

	txHash := msgTx.TxHash()
	blkHash, _ := chainhash.NewHashFromStr(cypherBlock.Hash)
	merkleProofs := btcrelaying.BuildMerkleProof(txHashes, &txHash)
	btcProof := btcrelaying.BTCProof{
		MerkleProofs: merkleProofs,
		BTCTx:        msgTx,
		BlockHash:    blkHash,
	}
	btcProofBytes, _ := json.Marshal(btcProof)
	btcProofStr := base64.StdEncoding.EncodeToString(btcProofBytes)

	// verify btc proof
	decodedProof, err := btcrelaying.ParseBTCProofFromB64EncodeStr(btcProofStr)
	require.Equal(t, nil, err)

	merkleRoot, _ := chainhash.NewHashFromStr(cypherBlock.MerkleRoot)
	isValid := btcrelaying.Verify(merkleRoot, decodedProof.MerkleProofs, &txHash)
	require.Equal(t, true, isValid)

	// run stimulate eth and deploy contract
	sim := runSimEth()
	_, _, btcContract, err := btcrelay.DeployBtcrelay(auth, sim)
	require.Equal(t, nil, err)

	sim.Commit()

	var merkleRootByte32 [32]byte
	var merkleProof []btcrelay.BTCProofMerkleProof
	copy(merkleRootByte32[:], merkleRoot.CloneBytes())

	for _, v := range decodedProof.MerkleProofs {
		var a [32]byte
		copy(a[:], v.ProofHash.CloneBytes())
		merkleProof = append(merkleProof, btcrelay.BTCProofMerkleProof{a, v.IsLeft})
	}
	//verify tx hash
	result, err := btcContract.Verify(nil, merkleRootByte32, merkleProof, txHash)
	require.Equal(t, nil, err)
	require.Equal(t, true, result)

	// input wrong txHash
	txHash[0]++
	result, err = btcContract.Verify(nil, merkleRootByte32, merkleProof, txHash)
	require.Equal(t, nil, err)
	require.Equal(t, false, result)
}


func runSimEth(accs ...ec.Address) *backends.SimulatedBackend {
	genesisAcc = loadAccount()
	auth = bind.NewKeyedTransactor(genesisAcc.PrivateKey)
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("1000000000000000000000000000000", 10) // 1E30 wei
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	for _, acc := range accs {
		alloc[acc] = core.GenesisAccount{Balance: balance}
	}
	sim := backends.NewSimulatedBackend(alloc, 8000000)
	return sim
}

func loadAccount() *account {
	key, err := crypto.LoadECDSA("genesisKey.hex")
	if err != nil {
		return newAccount()
	}
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func newAccount() *account {
	key, _ := crypto.GenerateKey()
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}
