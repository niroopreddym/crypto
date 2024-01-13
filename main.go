package main

import (
	"fmt"

	"github.com/niroopreddym/blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", string(block.PrevHash))
		fmt.Println("Data : ", string(block.Data))
		fmt.Printf("Hash: %x\n", string(block.Hash))

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %v\n", pow.Validate())
		fmt.Println()
	}
}
