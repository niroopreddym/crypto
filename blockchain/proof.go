package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/niroopreddym/blockchain/common"
)

//proof of work based alogorithm

//take any block

//create a counter (nounce) which starts at 0

//create a hash of data plus the counter

//check the hash to see it it meets requirements

//requirements:

// 1. the first few bytes must contain 0s

const difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int //this is the number that is helping us achieve our requirement
}

// this ctor takes in a pointer to block and craetes a poinetr to proof of work
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//difficlty nounce can be any arbtrary value and 256 is the sha256 encrypted data length
	target.Lsh(target, uint(256-difficulty))
	pow := &ProofOfWork{
		Block:  b,
		Target: target,
	}

	return pow
}

func (pow *ProofOfWork) InitData(nounce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		common.ToHex(int64(nounce)),
		common.ToHex(int64(difficulty))}, []byte{})

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nounce := 0

	for nounce < math.MaxInt64 {
		data := pow.InitData(nounce)
		hash = sha256.Sum256(data)
		fmt.Printf("%x\n", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nounce++
		}
	}

	fmt.Println()
	return nounce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nounce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}
