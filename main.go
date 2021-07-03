package main

import (
	"fmt"
	"github.com/Wh0rigin/Sake33Coin/functions"
	"github.com/Wh0rigin/Sake33Coin/structs"
)

func main() {
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

	//密钥
	key, _ := functions.MakeNewKey("he21412412125fwwggegwegwgewgwgwegwegewgewgewwdsvsshgvdvjykktykdsgsx")

	sakecoin := structs.NewChain()
	sakecoin.SetLevel(1)
	t1 := structs.NewTransaction("add1", "add2", "10")
	t1.Sign(key)
	//t1.amount = "20" 篡改数据会被检验
	fmt.Println(t1.IsValid(key))
	t2 := structs.NewTransaction("add2", "add1", "5")
	sakecoin.AddTransaction(t1, key)
	sakecoin.AddTransaction(t2, key)
	fmt.Println(*sakecoin)

	sakecoin.MineTransactionPool("add3", key)
	fmt.Println(*sakecoin)
}
