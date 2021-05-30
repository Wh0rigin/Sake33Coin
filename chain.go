package main

import(
	"strconv"
	"fmt"
)

type Chain struct{
	chain []Block
	transactionPool []Transaction
	mineReward int
	level int
}

func NewChain() *Chain {
	chain := &Chain{}
	chain.chain = []Block{chain.bigBang()}
	chain.level = 1
	chain.mineReward = 50
	chain.transactionPool = []Transaction{}
	return chain
}

func (this *Chain)bigBang() Block {
	genesisBlock := NewBlock([]Transaction{NewTransaction("先祖区块","先祖区块","")},"")
	return *genesisBlock
}

func (this *Chain)getlastBlock() Block {
	return this.chain[len(this.chain) - 1]
}

func (this *Chain)addBlockToChain(newBlock *Block,key *GKey){
	newBlock.previousHash = this.getlastBlock().hash
	//newBlock.hash = newBlock.computeHash()
	newBlock.mine(this.level,key)
	this.chain = append(this.chain,*newBlock)
}

func (this *Chain)mineTransactionPool(minerAddress string,key *GKey){
	//发放矿工奖励
	transaction := NewTransaction("",minerAddress,strconv.Itoa(this.mineReward))
	this.transactionPool = append(this.transactionPool,transaction)
	//挖矿
	newBlock := NewBlock(this.transactionPool,this.getlastBlock().hash)
	newBlock.mine(this.level,key)
	//添加区块到区块链
	this.chain = append(this.chain,*newBlock)
	this.transactionPool = []Transaction{}
}

func (this *Chain)validataChain(key *GKey) bool {
	if 1 == len(this.chain){
		if this.chain[0].hash != this.chain[0].computeHash() {
			return false
		}
		return true
	}
	for i:= 1;i < len(this.chain);i++{
		blockToValidata := this.chain[i]
		if !blockToValidata.validataBlockTransactions(key) {
			fmt.Println("发现非法交易")
			return false
		}
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

func (this *Chain)SetLevel(_level int){
	this.level = _level
}

func (this *Chain)addTransaction(transaction Transaction,key *GKey){
	if !transaction.isValid(key){
		fmt.Println("invalid transaction")
		return
	}
	fmt.Println("valid transaction")
	this.transactionPool = append(this.transactionPool,transaction)
}
