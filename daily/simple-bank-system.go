package main

import "math"

type Bank struct {
	Accounts []int64
}

func Constructor(balance []int64) Bank {
	var acc = []int64{0}
	for _, b := range balance {
		acc = append(acc, b)
	}
	return Bank{
		Accounts: acc,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > 0 && account2 > 0 && account1 <= len(this.Accounts) && account2 <= len(this.Accounts) {
		if this.Accounts[account1] >= money && math.MaxInt64-money > this.Accounts[account2] {
			this.Accounts[account1] = this.Accounts[account1] - money
			this.Accounts[account2] = this.Accounts[account2] + money
			return true
		}
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > 0 && account <= len(this.Accounts) && math.MaxInt64-money > this.Accounts[account] {
		this.Accounts[account] = this.Accounts[account] + money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > 0 && account <= len(this.Accounts) && this.Accounts[account] >= money {
		this.Accounts[account] = this.Accounts[account] - money
		return true
	}
	return false
}
