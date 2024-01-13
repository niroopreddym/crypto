package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/niroopreddym/blockchain/common"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nounce   int
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func (b *Block) Serialize() []byte {
	buff := bytes.Buffer{}
	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(b)
	common.PanicErr(err)

	return buff.Bytes()
}

func DeSerialize(data []byte) *Block {
	block := Block{}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	common.PanicErr(err)
	return &block
}

// ctor
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}

	pow := NewProof(block)
	nounce, hash := pow.Run()
	block.Hash = hash
	block.Nounce = nounce

	return block
}
