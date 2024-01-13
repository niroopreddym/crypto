package main

import (
	"fmt"

	"github.com/niroopreddym/blockchain/common"
)

func main() {
	chain := common.InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", string(block.PrevHash))
		fmt.Println("Data : ", string(block.Data))
		fmt.Printf("Hash: %x\n", string(block.Hash))

		fmt.Println()
	}
}
