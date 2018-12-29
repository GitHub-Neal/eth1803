package main

import (
	"awesomeProject/M3/day2/bkc/c02-bkc/BLC"
	"fmt"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1,
		bc.Blocks[len(bc.Blocks) - 1].Hash,
		[]byte("a send 100 btc to b"))
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1,
		bc.Blocks[len(bc.Blocks) - 1].Hash,
		[]byte("b send 100 btc to c"))

	for _,block:=range bc.Blocks{
		fmt.Printf("block:%v\n",block)
	}
}
