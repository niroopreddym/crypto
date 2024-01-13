package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/niroopreddym/blockchain/common"
	"github.com/niroopreddym/blockchain/database"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (chain *BlockChain) AddBlock(data string) {
	lastHash := []byte{}
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		common.PanicErr(err)
		lastHash, err = item.ValueCopy(nil)
		return err
	})

	common.PanicErr(err)
	newblock := CreateBlock(data, lastHash)
	chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newblock.Hash, newblock.Serialize())
		common.PanicErr(err)
		err = txn.Set([]byte("lh"), newblock.Hash)
		chain.LastHash = newblock.Hash
		return err
	})
	common.PanicErr(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{
		CurrentHash: chain.LastHash,
		Database:    chain.Database,
	}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	block := &Block{}
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		encodedBlock, err := item.ValueCopy(nil)
		block = DeSerialize(encodedBlock)
		return err
	})

	common.PanicErr(err)

	iter.CurrentHash = block.PrevHash
	return block
}

// genisys block -- also called as the zero block for any block chain
func Genesis() *Block {
	return CreateBlock("Genysis", []byte{})
}

func InitBlockChain() *BlockChain {
	db := database.NewDB()
	lastHash := []byte{}
	err := db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err != nil {
			fmt.Println("no blockchain found")
			genisys := Genesis()
			fmt.Println("Genisys Proved")
			err = txn.Set(genisys.Hash, genisys.Serialize())
			common.PanicErr(err)
			err = txn.Set([]byte("lh"), genisys.Hash)
			lastHash = genisys.Hash
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			common.PanicErr(err)
			lastHash, err = item.ValueCopy(nil)
			return err
		}
	})

	common.PanicErr(err)

	return &BlockChain{
		LastHash: lastHash,
		Database: db,
	}
}
