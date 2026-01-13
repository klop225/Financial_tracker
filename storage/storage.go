package storage

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func ReadBalance(money float64, b bool) {
	data, _ := os.ReadFile("balance")
	balance, _ := strconv.ParseFloat(string(data), 64)

	if b {
		WriteBalance(balance + money)
		WriteHistory("Пополнил", money, balance+money)
		return
	} else if money < balance {
		WriteBalance(balance - money)
		WriteHistory("Снял", money, balance-money)
		return
	}
	fmt.Println("столько на баласне нету")
}

func WriteBalance(balance float64) {
	file, _ := os.Create("balance")

	defer file.Close()

	fmt.Fprintf(file, "%.2f", balance)
}

func WriteHistory(operation string, money, balance float64) {
	file, _ := os.OpenFile("history", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	time := time.Now().Format("2006-01-02 15:04:05")

	fmt.Fprintf(file, "[%s] %s %.2f баланс сейчас %.2f\n", time, operation, money, balance)
}
