package main

import(
	"crypto/sha256"
	"fmt"
)

type Block struct{
	data string
	previousHash string
	hash string
}

func NewBlock(_data,_previousHash string)*Block{
	block := &Block{
		data : _data,
		previousHash : _previousHash,
	}
	block.hash = block.computeHash()
	return block
}
func (this *Block)computeHash() string {
	shaWr := sha256.New()
	shaWr.Write([]byte(this.data))
	return fmt.Sprintf("%x",shaWr.Sum([]byte(this.previousHash)))
}
