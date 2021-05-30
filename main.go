package main

import(
	"fmt"
)


func main(){
	//block := NewBlock("转账10元","214")
	//fmt.Println(*block)

	//SakeChain := NewChain()
	//fmt.Println(*SakeChain)
	//SakeChain.SetLevel(10)
	//SakeChain.addBlockToChain(block)
	//fmt.Println(*SakeChain)

	//if SakeChain.validataChain() {
	//	fmt.Println("数据没有被篡改")
	//}else{
	//	fmt.Println("数据被篡改")
	//}

	//SakeChain.chain[1].data = "转账一个亿"
	//fmt.Println(*SakeChain)
     	//if SakeChain.validataChain() {
        //        fmt.Println("数据没有被篡改")
        //}else{
        //        fmt.Println("数据被篡改")
        //}

	sakecoin := NewChain()
	sakecoin.SetLevel(3)
	t1 := NewTransaction("add1","add2","10")
	t2 := NewTransaction("add2","add1","5")
	sakecoin.addTransaction(t1)
	sakecoin.addTransaction(t2)
	fmt.Println(*sakecoin)

	sakecoin.mineTransactionPool("add3")
	fmt.Println(*sakecoin)
}
