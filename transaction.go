package main

import(
	
)

type Transaction struct{
	from string      `json:from`
	to string        `json:to`
	amount string	 `json:amount`
	signature string `json:signature`
}

func NewTransaction(_from,_to,_amount string) Transaction {
	transaction := &Transaction{
		from : _from,
		to : _to,
		amount : _amount,
	}
	return *transaction
}

func (this *Transaction)sign(key *GKey){
	this.signature,_ = key.Sign([]byte(this.computeHash()))
}

func (this *Transaction)isValid(key *GKey) bool {
	if this.from == ""{
		return true
	}

	res , _ := Verify([]byte(this.computeHash()),this.signature,&key.PublicKey)
	return res
}

func (this *Transaction)computeHash() string {
	return GetSHA256HashCode(this.from + this.to + this.amount)
}
