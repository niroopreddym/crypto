package database

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/niroopreddym/blockchain/common"
)

var (
	dbPath = "./tmp/blocks"
)

type TxOutput struct {
	Value  int
	PubKey string
}
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func NewDB() *badger.DB {
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	common.PanicErr(err)
	return db
}
