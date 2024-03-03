package main

import (
	"fmt"
)

// Block represents a data block.
type Block struct {
	Data string
}

// DisplayAllBlocks prints the data of all blocks in the given slice.
func DisplayAllBlocks(blocks []Block) {
	for _, block := range blocks {
		fmt.Println(block.Data)
	}
}

// NewBlock creates a new block with the given data.
func NewBlock(data string) Block {
	return Block{Data: data}
}

// ModifyBlock modifies the data of the given block.
func ModifyBlock(block *Block, newData string) {
	block.Data = newData
}

func main() {
	// Example usage:
	blocks := []Block{
		NewBlock("Block 1"),
		NewBlock("Block 2"),
		NewBlock("Block 3"),
	}

	DisplayAllBlocks(blocks)

	// Modify the second block
	ModifyBlock(&blocks[1], "Modified Block 2")

	DisplayAllBlocks(blocks)
}
