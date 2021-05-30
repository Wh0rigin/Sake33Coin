package main

import(
)

type Transaction struct{
	from string      `json:from`
	to string        `json:to`
	amount string	 `json:amount`
}

func NewTransaction(_from,_to,_amount string) Transaction {
	transaction := &Transaction{
		from : _from,
		to : _to,
		amount : _amount,
	}
	return *transaction
}
