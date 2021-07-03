package structs

import (
	"fmt"
	"github.com/Wh0rigin/Sake33Coin/functions"
	"strconv"
	"time"
)

type Block struct {
	data         []Transaction
	previousHash string
	timestamp    string
	nonce        int
	hash         string
}

func NewBlock(_data []Transaction, _previousHash string) *Block {
	block := &Block{
		data:         _data,
		previousHash: _previousHash,
		timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		nonce:        1,
	}
	block.hash = block.computeHash()
	return block
}
func (this *Block) computeHash() string {
	return functions.GetSHA256HashCode(GetJsonarraytoString(this.data) + this.previousHash + strconv.Itoa(this.nonce) + this.timestamp)
}

func (this *Block) getAnswer(level int) string {
	answer := ""
	for i := 0; i < level; i++ {
		answer += "0"
	}
	return answer
}

func (this *Block) mine(level int, key *functions.GKey) {
	if !this.validataBlockTransactions(key) {
		return
	}
	for {
		this.hash = this.computeHash()
		if this.hash[0:level] != this.getAnswer(level) {
			this.nonce++
			this.hash = this.computeHash()
			fmt.Println("挖矿中:", this.hash)
		} else {
			break
		}
	}
	fmt.Println("挖矿获得:", this.hash)
}

func (this *Block) validataBlockTransactions(key *functions.GKey) bool {
	for _, transaction := range this.data {
		if false == transaction.IsValid(key) {
			fmt.Println("invalid transaction found in transactions:发现异常交易")
			return false
		}
	}
	return true
}
