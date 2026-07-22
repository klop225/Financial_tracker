package command

type Storage interface {
	Deposit(money float64) error
	Withdraw(money float64) error
	GetBalance() (float64, error)
	GetHistory() (string, error)
}
type CommandHandler struct {
	store Storage
}

func NewCommandHandler(store Storage) *CommandHandler {
	return &CommandHandler{store: store}
}

func (c *CommandHandler) Add(money float64) error {
	return c.store.Deposit(money)
}

func (c *CommandHandler) Take(money float64) error {
	return c.store.Withdraw(money)
}

func (c *CommandHandler) GetBalance() (float64, error) {
	return c.store.GetBalance()
}

func (c *CommandHandler) GetHistory() (string, error) {
	return c.store.GetHistory()
}
