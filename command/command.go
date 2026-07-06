package command

import "Financial_tracker/storage"

func Add(money float64) error {
	return storage.Deposit(money)
}

func Take(money float64) error {
	return storage.Withdraw(money)
}

func GetBalance() (float64, error) {
	return storage.GetBalance()
}

func GetHistory() (string, error) {
	return storage.GetHistory()
}
