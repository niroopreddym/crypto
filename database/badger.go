package database

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/niroopreddym/blockchain/common"
)

var (
	dbPath = "./tmp/blocks"
)

func NewDB() *badger.DB {
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	common.PanicErr(err)
	return db
}
