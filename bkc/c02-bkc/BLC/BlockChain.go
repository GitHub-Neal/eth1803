package BLC

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock()*BlockChain  {
	block:=CreateGenesisBlock("the init of the blockchain")
	return &BlockChain{[]*Block{block}}
}

func (bc *BlockChain)AddBlock(height int64,prevBlockHash []byte,data []byte)  {
	newBlock:=NewBlock(height,prevBlockHash,data)
	bc.Blocks=append(bc.Blocks,newBlock)
}