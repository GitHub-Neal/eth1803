package main

import (
	"github.com/boltdb/bolt"
	"fmt"
	"awesomeProject/M3/day2/bkc/c05-bc-db/BLC"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()
	bc.AddBlock([]byte("a send 100 btc to b"))
	bc.AddBlock([]byte("b send 100 btc to c"))

	bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if nil != b {
			hash := b.Get([]byte("l"))
			fmt.Printf("value : %x\n", hash)
			block := b.Get(hash)
			fmt.Printf("height : %d\n", BLC.DeserializeBlock(block).Height)
		} else {
			fmt.Printf("the bucket is nil\n")
		}
		return nil
	})
}
