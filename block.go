package main

import(
	"fmt"
	"strconv"
	"time"
)

type Block struct{
	data []Transaction
	previousHash string
	timestamp string
	nonce int
	hash string
}

func NewBlock(_data []Transaction,_previousHash string)*Block{
	block := &Block{
		data : _data,
		previousHash : _previousHash,
		timestamp : time.Now().Format("2006-01-02 15:04:05"),
		nonce :1,
	}
	block.hash = block.computeHash()
	return block
}
func (this *Block)computeHash() string {
	return GetSHA256HashCode(GetJsonarraytoString(this.data)+this.previousHash+strconv.Itoa(this.nonce)+this.timestamp)
}

func (this *Block)getAnswer(level int) string {
	answer := ""
	for i := 0;i < level;i++{
		answer += "0"
	}
	return answer
}

func (this *Block)mine(level int){
	for{
		this.hash = this.computeHash()
		if this.hash[0:level] != this.getAnswer(level){
			this.nonce++
			this.hash = this.computeHash()
			fmt.Println("挖矿中:",this.hash)
		}else{
			break
		}
	}
	fmt.Println("挖矿获得:",this.hash)
}
