package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	jsonPure := []byte(`{"Number":2,"Balance":200}`)
	var accountX Account
	err = json.Unmarshal(jsonPure, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)

	jsonPure2 := []byte(`{"n":2,"b":200}`)
	var accountX2 Account
	err = json.Unmarshal(jsonPure2, &accountX2)
	if err != nil {
		println(err)
	}
	println(accountX2.Balance)
}
