package main

import "fmt"

type account struct {
	name string
}

//  新建账号
func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

// 检查账号
func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
