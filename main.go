package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/niroopreddym/blockchain/blockchain"
	"github.com/niroopreddym/blockchain/common"
)

type commandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *commandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("add - block BLOCK_DATA - add block to the chain")
	fmt.Println("print - Prints the blocks in chain ")
}

func (cli *commandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *commandLine) run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addBlockCmdArgs := addBlockCmd.String("block", "", "Block Data")

	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		common.PanicErr(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		common.PanicErr(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockCmdArgs == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}

		cli.addBlock(*addBlockCmdArgs)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *commandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added New Block!")
}

func (cli *commandLine) printChain() {

	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()
		fmt.Printf("Prev Hash: %x\n", string(block.PrevHash))
		fmt.Println("Data : ", string(block.Data))
		fmt.Printf("Hash: %x\n", string(block.Hash))

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %v\n", pow.Validate())
		fmt.Printf("Nounce: %v\n", pow.Block.Nounce)
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()
	cli := commandLine{
		chain,
	}

	cli.run()
}
