package main

import "fmt"

type wallet struct {
	balance int // 余额
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

// 余额查询
func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

// 支付
func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not Sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}
