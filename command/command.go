package command

import "Financial_tracker/storage"

type CommandHadler struct {
	store *storage.File
}

func NewCommandHadler(store *storage.File) *CommandHadler {
	return &CommandHadler{store: store}
}

func (c *CommandHadler) Add(money float64) error {
	return c.store.Deposit(money)
}

func (c *CommandHadler) Take(money float64) error {
	return c.store.Withdraw(money)
}

func (c *CommandHadler) GetBalance() (float64, error) {
	return c.store.GetBalance()
}

func (c *CommandHadler) GetHistory() (string, error) {
	return c.store.GetHistory()
}
