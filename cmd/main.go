package main

import (
	"fmt"

	"github.com/JovidYnwa/wallet/pkg/wallet"
)

func main() {

	svc := &wallet.Service{}
	//wallet.RegisterAccount(svc, "+992907013487")
	//svc.RegisterAccount("+992907013487")
	//svc.Deposit(1, 10)
	account, err := svc.RegisterAccount("+992907013487")
	if err != nil {
		fmt.Println(*svc)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err := svc.Pay(account.ID, 88, "some_category")
	findAccount, err := svc.FindAccountByID(2)
	findPayment, err := svc.FindPaymentByID("1")
	fmt.Println(account)
	fmt.Println(account.Balance, account.ID)
	fmt.Println(payment)
	fmt.Println(findAccount)
	fmt.Println(findPayment)

}
