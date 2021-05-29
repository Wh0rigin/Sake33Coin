package main

type Chain struct{
	chain []Block
}

func NewChain() *Chain {
	chain := &Chain{}
	chain.chain = []Block{chain.bigBang()}
	return chain
}

func (this *Chain)bigBang() Block {
	genesisBlock := NewBlock("我是祖先区块","")
	return *genesisBlock
}

func (this *Chain)getlastBlock() Block {
	return this.chain[len(this.chain) - 1]
}

func (this *Chain)addBlockToChain(newBlock *Block){
	newBlock.previousHash = this.getlastBlock().hash
	newBlock.hash = newBlock.computeHash()
	this.chain = append(this.chain,*newBlock)
}

func (this *Chain)validataChain() bool {
	if 1 == len(this.chain){
		if this.chain[0].hash != this.chain[0].computeHash() {
			return false
		}
		return true
	}
	for i:= 1;i < len(this.chain);i++{
		blockToValidata := this.chain[i]
		if blockToValidata.hash != blockToValidata.computeHash() {
			return false
		}
		previousBlock := this.chain[i-1]
		if blockToValidata.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
