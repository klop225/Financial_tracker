package storage

import (
	"os"
	"testing"
)

func TestDeposit(t *testing.T) {
	testStore := NewFile("balanceTest.json", "historyTest")

	defer func() {
		os.Remove("balanceTest.json")
		os.Remove("historyTest")
	}()

	err := testStore.Deposit(100.50)
	if err != nil {
		t.Errorf("Deposit вернул ошибку: %v", err)
	}

	bal, err := testStore.ReadBalance()
	if err != nil {
		t.Errorf("Не удалось прочитать баланс: %v", err)
	}

	if bal.Amount != 100.50 {
		t.Errorf("Ожидался баланс 100.50, получено %.2f", bal.Amount)
	}
}

func TestWithdraw(t *testing.T) {
	testStore := NewFile("balanceTest.json", "historyTest")

	defer func() {
		os.Remove("balanceTest.json")
		os.Remove("historyTest")
	}()

	err := testStore.Deposit(200)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = testStore.Withdraw(50)
	if err != nil {
		t.Errorf("Withdraw вернул ошибку: %v", err)
	}

	bal, err := testStore.ReadBalance()
	if err != nil {
		t.Errorf("Не удалось прочитать баланс: %v", err)
	}
	if bal.Amount != 150 {
		t.Errorf("Ожидался баланс 150.00, получено %.2f", bal.Amount)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	testStore := NewFile("balanceTest.json", "historyTest")

	defer func() {
		os.Remove("balanceTest.json")
		os.Remove("historyTest")
	}()

	err := testStore.Deposit(100)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = testStore.Withdraw(200)
	if err == nil {
		t.Error("Ожидалась ошибка о недостатке средств, но ее не было")
	}

	bal, _ := testStore.ReadBalance()
	if bal.Amount != 100 {
		t.Errorf("Баланс изменился: ожидалось 100.00, получено %.2f", bal.Amount)
	}
}

func TestDepositNegative(t *testing.T) {
	testStore := NewFile("balanceTest.json", "historyTest")

	defer func() {
		os.Remove("balanceTest.json")
		os.Remove("historyTest")
	}()

	err := testStore.Deposit(-50)
	if err == nil {
		t.Error("Ожидалась ошибка при отрицательной сумме")
	}
}

func TestWithdrawNegative(t *testing.T) {
	testStore := NewFile("balanceTest.json", "historyTest")

	defer func() {
		os.Remove("balanceTest.json")
		os.Remove("historyTest")
	}()

	err := testStore.Deposit(100)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = testStore.Withdraw(-50)
	if err == nil {
		t.Error("Ожидалась ошибка при отрицательной сумме")
	}
}

func TestGetHistory(t *testing.T) {
	testStore := NewFile("balanceTest.json", "historyTest")

	defer func() {
		os.Remove("balanceTest.json")
		os.Remove("historyTest")
	}()

	history, err := testStore.GetHistory()
	if err != nil {
		t.Errorf("GetHistory вернул ошибку: %v", err)
	}

	if history != "история пуста\n" {
		t.Errorf("Ожидалась пустая история, получено: %s", history)
	}

	err = testStore.Deposit(30)
	if err != nil {
		t.Errorf("Не удалось пополнить баланс: %v", err)
	}

	err = testStore.Withdraw(20)
	if err != nil {
		t.Errorf("Withdraw вернул ошибку: %v", err)
	}

	h, err := testStore.GetHistory()
	if err != nil {
		t.Errorf("GetHistory вернул ошибку: %v", err)
	}

	if len(h) == 0 {
		t.Error("История не должна быть пустой")
	}
}
