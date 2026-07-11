package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type File struct {
	balanceFile string
	historyFile string
}
type Balance struct {
	Amount float64 `json:"amount"`
}

func NewFile(balFile, HistFile string) *File {
	return &File{
		balanceFile: balFile,
		historyFile: HistFile,
	}
}

func (f *File) ReadBalance() (Balance, error) {
	data, err := os.ReadFile(f.balanceFile)
	if err != nil {
		if os.IsNotExist(err) {
			return Balance{Amount: 0}, nil
		}
		return Balance{}, fmt.Errorf("не удалось прочитать баланс: %w", err)
	}

	var bal Balance
	if err := json.Unmarshal(data, &bal); err != nil {
		return Balance{}, fmt.Errorf("файл баланса повреждён: %w", err)
	}

	return bal, nil
}

func (f *File) WriteBalance(balance float64) error {
	bal := Balance{Amount: balance}
	data, err := json.Marshal(bal)
	if err != nil {
		return fmt.Errorf("не удалось сформировать json: %w", err)
	}

	if err := os.WriteFile(f.balanceFile, data, 0644); err != nil {
		return fmt.Errorf("не удалось записать баланс: %w", err)
	}
	return nil
}

func (f *File) writeHistory(operation string, money, balance float64) error {
	file, err := os.OpenFile(f.historyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть историю: %w", err)
	}
	defer file.Close()

	t := time.Now().Format("2006-01-02 15:04:05")
	_, err = fmt.Fprintf(file, "[%s] %s %.2f баланс сейчас %.2f\n", t, operation, money, balance)
	if err != nil {
		return fmt.Errorf("не удалось записать историю: %w", err)
	}
	return nil
}

func (f *File) Deposit(money float64) error {
	if money <= 0 {
		return fmt.Errorf("некорректная сумма пополнения: %.2f, сумма должна быть больше нуля", money)
	}

	bal, err := f.ReadBalance()
	if err != nil {
		return err
	}
	newAmount := bal.Amount + money
	if err := f.WriteBalance(newAmount); err != nil {
		return err
	}
	return f.writeHistory("Пополнил", money, newAmount)
}

func (f *File) Withdraw(money float64) error {
	if money <= 0 {
		return fmt.Errorf("некорректная сумма снятия: %.2f, сумма должна быть больше нуля", money)
	}

	bal, err := f.ReadBalance()
	if err != nil {
		return err
	}

	if money > bal.Amount {
		return fmt.Errorf("недостаточно средств: на балансе %.2f, запрошено %.2f", bal.Amount, money)
	}

	newAmount := bal.Amount - money
	if err := f.WriteBalance(newAmount); err != nil {
		return err
	}
	return f.writeHistory("Снял", money, newAmount)
}

func (f *File) GetBalance() (float64, error) {
	bal, err := f.ReadBalance()
	if err != nil {
		return 0, err
	}
	return bal.Amount, nil
}

func (f *File) GetHistory() (string, error) {
	data, err := os.ReadFile(f.historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "история пуста\n", nil
		}
		return "", fmt.Errorf("не удалось прочитать историю: %w", err)
	}
	return string(data), nil
}
