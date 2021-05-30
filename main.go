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
	

	//密钥
	key,_ := MakeNewKey("he21412412125fwwggegwegwgewgwgwegwegewgewgewwdsvsshgvdvjykktykdsgsx")


	sakecoin := NewChain()
	sakecoin.SetLevel(1)
	t1 := NewTransaction("add1","add2","10")
	t1.sign(key)
	//t1.amount = "20" 篡改数据会被检验
	fmt.Println(t1.isValid(key))
	t2 := NewTransaction("add2","add1","5")
	sakecoin.addTransaction(t1,key)
	sakecoin.addTransaction(t2,key)
	fmt.Println(*sakecoin)

	sakecoin.mineTransactionPool("add3",key)
	fmt.Println(*sakecoin)
}
