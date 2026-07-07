package storage

import (
	"fmt"
	"os"
	"testing"
)

const (
	balanceFIle = "balance.json"
	historyFIle = "history"
)

func TestDeposit(t *testing.T) {

	defer func() {
		os.Remove(balanceFIle)
		os.Remove(historyFIle)
	}()

	err := Deposit(100.50)
	if err != nil {
		t.Errorf("Deposit вернул ошибку: %v", err)
	}

	bal, err := ReadBalance()
	if err != nil {
		t.Errorf("Не удалось прочитать баланс: %v", err)
	}

	if bal.Amount != 100.50 {
		t.Errorf("Ожидался баланс 100.50, получено %.2f", bal.Amount)
	}
}

func TestWithdraw(t *testing.T) {

	defer func() {
		os.Remove(balanceFile)
		os.Remove(historyFile)
	}()

	err := Deposit(200)
	if err != nil {
		fmt.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = Withdraw(50)
	if err != nil {
		fmt.Errorf("Withdraw вернул ошибку: %v", err)
	}

	bal, err := ReadBalance()
	if err != nil {
		fmt.Errorf("Не удалось прочитать баланс: %v", err)
	}
	if bal.Amount != 150 {
		fmt.Errorf("Ожидался баланс 150.00, получено %.2f", bal.Amount)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {

	defer func() {
		os.Remove(balanceFile)
		os.Remove(historyFile)
	}()

	err := Deposit(100)
	if err != nil {
		fmt.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = Withdraw(200)
	if err == nil {
		t.Error("Ожидалась ошибка о недостатке средств, но ее не было")
	}

	bal, _ := ReadBalance()
	if bal.Amount != 100 {
		t.Errorf("Баланс изменился: ожидалось 100.00, получено %.2f", bal.Amount)
	}
}

func TestDepositNegative(t *testing.T) {

	defer func() {
		os.Remove(balanceFile)
		os.Remove(historyFile)
	}()

	err := Deposit(-50)
	if err == nil {
		t.Error("Ожидалась ошибка при отрицательной сумме")
	}
}
