package storage

import (
	"os"
	"testing"
)

const (
	balanceFileTest = "balanceTest.json"
	historyFileTest = "historyTest"
)

func init() {
	balanceFile = balanceFileTest
	historyFile = historyFileTest
}

func TestDeposit(t *testing.T) {

	defer func() {
		os.Remove(balanceFileTest)
		os.Remove(historyFileTest)
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
		os.Remove(balanceFileTest)
		os.Remove(historyFileTest)
	}()

	err := Deposit(200)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = Withdraw(50)
	if err != nil {
		t.Errorf("Withdraw вернул ошибку: %v", err)
	}

	bal, err := ReadBalance()
	if err != nil {
		t.Errorf("Не удалось прочитать баланс: %v", err)
	}
	if bal.Amount != 150 {
		t.Errorf("Ожидался баланс 150.00, получено %.2f", bal.Amount)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {

	defer func() {
		os.Remove(balanceFileTest)
		os.Remove(historyFileTest)
	}()

	err := Deposit(100)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
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
		os.Remove(balanceFileTest)
		os.Remove(historyFileTest)
	}()

	err := Deposit(-50)
	if err == nil {
		t.Error("Ожидалась ошибка при отрицательной сумме")
	}
}

func TestWithdrawNegative(t *testing.T) {

	defer func() {
		os.Remove(balanceFileTest)
		os.Remove(historyFileTest)
	}()

	err := Deposit(100)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = Withdraw(-50)
	if err == nil {
		t.Error("Ожидалась ошибка при отрицательной сумме")
	}
}

func TestGetHistory(t *testing.T) {

	defer func() {
		os.Remove(balanceFileTest)
		os.Remove(historyFileTest)
	}()

	history, err := GetHistory()
	if err != nil {
		t.Errorf("GetHistory вернул ошибку: %v", err)
	}

	if history != "история пуста\n" {
		t.Errorf("Ожидалась пустая история, получено: %s", history)
	}

	err = Deposit(30)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = Withdraw(20)
	if err != nil {
		t.Errorf("Withdraw вернул ошибку: %v", err)
	}

	h, err := GetHistory()
	if err != nil {
		t.Errorf("GetHistory вернул ошибку: %v", err)
	}

	if len(h) == 0 {
		t.Error("История не должна быть пустой")
	}
}
