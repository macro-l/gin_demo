package main

import "fmt"

// 钱包的结构
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

// 钱包门面
func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("starting create account")

	walletFacade := &walletFacade{
		account:      newAccount(accountID), // 根据 accountID 检查并获取一个账号
		securityCode: newSecurityCode(code), // 根据 code 检查安全code
		wallet:       newWallet(),           // 获取钱包
		notification: &notification{},       // 获取通知队列
		ledger:       &ledger{},             // 消费列表
	}

	fmt.Println("Account created")
	return walletFacade
}

// 充值
func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	// 更新余额
	w.wallet.creditBalance(amount)
	// 发送余额更新通知
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil

}

// 支付
func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	// 余额查询
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
