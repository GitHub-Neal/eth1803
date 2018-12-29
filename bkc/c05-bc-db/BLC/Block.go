package BLC

import (
	"time"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	TimeStamp int64
	Hash []byte
	PrevBlockHash []byte
	Data []byte
	Height int64
	Nonce int64
}

func (b *Block) SetHash() {
	timeStampBytes:=IntToHex(b.TimeStamp)
	heightBytes:=IntToHex(b.Height)
	blockBytes:=bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	},[]byte{})
	hash:=sha256.Sum256(blockBytes)
	b.Hash=hash[:]
}


func NewBlock(height int64,prevBlockHash []byte,data []byte)*Block{
	var block Block
	block=Block{
		TimeStamp:time.Now().Unix(),
		Hash:nil,
		PrevBlockHash:prevBlockHash,
		Data:data,
		Height:height,
	}
	//block.SetHash()
	pow:=NewProofWork(&block)
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.Nonce=nonce
	return &block
}

func CreateGenesisBlock(data string)*Block{
	return NewBlock(1,nil,[]byte(data))
}

func (block *Block)Serialize()[]byte  {
	var result bytes.Buffer
	encoder:=gob.NewEncoder(&result)
	if err:=encoder.Encode(block);nil!=err{
		log.Panicf("serialize the block to []byte failed! %v\n",err)
	}
	return result.Bytes()
}

func DeserializeBlock(blockBytes []byte)*Block{
	var block Block
	decoder:=gob.NewDecoder(bytes.NewReader(blockBytes))
	if err:=decoder.Decode(&block);nil!=err{
		log.Panicf("deserialize the []byte to block failed! %v\n", err)
	}
	return &block
}