package main

type Account struct {
	balance int
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

func (a *Account) simulate(value int) int {
	a.balance += value
	println(a.balance)
	return a.balance
}

func main() {
	account := Account{
		balance: 100,
	}
	account.simulate(200)
	println(account.balance)
}
