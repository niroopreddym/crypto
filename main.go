package main

import (
	"os"

	"github.com/niroopreddym/blockchain/blockchain"
	"github.com/niroopreddym/blockchain/cmd"
)

func main() {
	defer os.Exit(0)
	//-----------------block chain add blocks---------
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := cmd.CommandLine{
		Blockchain: chain,
	}

	cli.Run()
	//--------end--------block chain add blocks---------
}
