package BLC

import (
	"time"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	TimeStamp int64
	Hash []byte
	PrevBlockHash []byte
	Data []byte
	Height int64
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
	block.SetHash()
	return &block
}

func CreateGenesisBlock(data string)*Block{
	return NewBlock(1,nil,[]byte(data))
}