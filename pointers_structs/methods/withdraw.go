package main

import (
	"fmt"
)

type BankAccount struct {
  holder string
  balance int
}

func (account *BankAccount) withdraw(val int) {
	if account.balance >= val {
		account.balance -= val
	} else {
		fmt.Println("Insufficient Funds")
	}
}

func main() {
  acc := BankAccount{"James Smith", 100000}

  var amount int
  fmt.Scanln(&amount)

  acc.withdraw(amount)
  fmt.Println(acc)
}
