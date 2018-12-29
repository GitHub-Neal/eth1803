package BLC

import (
	"math/big"
	"crypto/sha256"
	"fmt"
	"bytes"
)

const targetBit=16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func NewProofWork(block *Block)*ProofOfWork  {
	target:=big.NewInt(1)
	target=target.Lsh(target,256-targetBit)
	return &ProofOfWork{block,target}
}

func (proofOfWork *ProofOfWork)Run()([]byte,int64)  {
	var nonce =0
	var hash [32]byte
	var hashInt big.Int
	for {
		dataBytes:=proofOfWork.prepareData(nonce)
		hash=sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		if proofOfWork.target.Cmp(&hashInt)==1{
			fmt.Printf("hash:%x\n",hash)
			break
		}
		nonce++
	}
	fmt.Printf("\n碰撞次数:%d\n",nonce)
	return hash[:],int64(nonce)
}

func (pow *ProofOfWork)prepareData(nonce int)[]byte{
	var data []byte
	data=bytes.Join([][]byte{
		IntToHex(pow.Block.Height),
		IntToHex(pow.Block.TimeStamp),
		pow.Block.PrevBlockHash,
		pow.Block.Data,
		IntToHex(int64(nonce)),
		IntToHex(targetBit),
	},[]byte{})
	return data
}