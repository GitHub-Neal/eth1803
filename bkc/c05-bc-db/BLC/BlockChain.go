package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const dbName  = "block.db"
const blockTableName ="blocks"
type BlockChain struct {
	DB *bolt.DB
	Tip []byte
}

func CreateBlockChainWithGenesisBlock()*BlockChain  {
	var blockHash []byte
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		fmt.Printf("b : %v\n", b)
		if nil == b {
			b, err := tx.CreateBucket([]byte(blockTableName))
			if nil != err {
				log.Panicf("create the bucket [%s] failed! %v\n", blockTableName, err)
			}
			genesisBlock := CreateGenesisBlock("the init of the blockchain")
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if nil != err {
				log.Panicf("insert the genesis block to db failed! %v\n", err)
			}
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if nil != err {
				log.Panicf("insert the latest block hash to db failed! %v\n", err)
			}
			blockHash = genesisBlock.Hash
			fmt.Printf("blockHash : %x\n", blockHash)
		}
		return nil
	})
	return &BlockChain{db, blockHash}
}

func (bc *BlockChain)AddBlock(data []byte)  {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			blockBytes := b.Get(bc.Tip)
			fmt.Printf("tip : %v\n", bc.Tip)
			latest_block := DeserializeBlock(blockBytes)
			newBlock := NewBlock(latest_block.Height + 1, latest_block.Hash, data)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("insert the new block to db failed! %v\n", err)
			}
			err = b.Put([]byte("l"), newBlock.Hash)
			if nil != err {
				log.Panicf("put the latest block hash to db failed! %v\n", err)
			}
			bc.Tip = newBlock.Hash
		}
		return nil
})
	if nil!=err{
	log.Panicf("update the db failed! %v\n",err)
	}
}