package main

import (
	"awesomeProject/M3/day2/bkc/c02-bkc/BLC"
	"fmt"
)

func main() {
	block:=BLC.NewBlock(1,nil,[]byte("this is the first block"))
	fmt.Printf("block-hash:%x\n",block.Hash)
}
