package structs

import (
	"encoding/json"
	"fmt"
	"github.com/Wh0rigin/Sake33Coin/functions"
)

type Transaction struct {
	from      string `json:from`
	to        string `json:to`
	amount    string `json:amount`
	signature string `json:signature`
}

func NewTransaction(_from, _to, _amount string) Transaction {
	transaction := &Transaction{
		from:   _from,
		to:     _to,
		amount: _amount,
	}
	return *transaction
}

func (this *Transaction) Sign(key *functions.GKey) {
	this.signature, _ = key.Sign([]byte(this.computeHash()))
}

func (this *Transaction) IsValid(key *functions.GKey) bool {
	if this.from == "" {
		return true
	}

	res, _ := functions.Verify([]byte(this.computeHash()), this.signature, &key.PublicKey)
	return res
}

func (this *Transaction) computeHash() string {
	return functions.GetSHA256HashCode(this.from + this.to + this.amount)
}

func GetJsonarraytoString(arr []Transaction) (res string) {
	if len(arr) == 0 {
		return
	}
	for _, i := range arr {
		jsonval, err := json.Marshal(i)
		if err != nil {
			fmt.Println("账单格式错误", err)
		}
		res += string(jsonval) + ","
	}
	res = res[:len(res)-1] //去最后一个,
	return
}
