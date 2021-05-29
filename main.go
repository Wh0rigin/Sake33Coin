package main

import(
	"fmt"
)


func main(){
	block := NewBlock("转账10元","214")
	fmt.Println(*block)

	SakeChain := NewChain()
	fmt.Println(*SakeChain)
	
	SakeChain.addBlockToChain(block)
	fmt.Println(*SakeChain)

	if SakeChain.validataChain() {
		fmt.Println("数据没有被篡改")
	}else{
		fmt.Println("数据被篡改")
	}

	SakeChain.chain[1].data = "转账一个亿"
	fmt.Println(*SakeChain)
     	if SakeChain.validataChain() {
                fmt.Println("数据没有被篡改")
        }else{
                fmt.Println("数据被篡改")
        }
}
