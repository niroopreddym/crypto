package cmd

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/niroopreddym/blockchain/blockchain"
	"github.com/niroopreddym/blockchain/common"
	"github.com/niroopreddym/blockchain/wallet"
)

type CommandLine struct {
	Blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("createblockchain - block BLOCK_DATA - add block to the chain")
	fmt.Println("printchain - Prints the blocks in chain ")
	fmt.Println("createwallet - Creates a new Wallet")
	fmt.Println("listaddresses - lists Address in wallet file")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	addBlockCmdArgs := addBlockCmd.String("block", "", "Block Data")

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)

	switch os.Args[1] {
	case "createblockchain":
		err := addBlockCmd.Parse(os.Args[2:])
		common.PanicErr(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		common.PanicErr(err)
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		common.PanicErr(err)
	case "listaddresses":
		err := listAddressCmd.Parse(os.Args[2:])
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

	if listAddressCmd.Parsed() {
		cli.listAddresses()
	}

	if createWalletCmd.Parsed() {
		cli.AddSingleWallet()
	}
}

func (cli *CommandLine) listAddresses() {
	wallets := wallet.CreateWallets()
	addresses := wallets.GetAllAddress()
	for address := range addresses {
		fmt.Println(address)
	}
}

func (cli *CommandLine) AddSingleWallet() {
	wallets := wallet.CreateWallets()
	address := wallets.AddWallet()
	wallets.SaveFile()
	fmt.Printf("New Address is: %s\n", address)
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("Added New Block!")
}

func (cli *CommandLine) printChain() {

	iter := cli.Blockchain.Iterator()

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
