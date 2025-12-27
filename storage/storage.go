package storage

import (
	"fmt"
	"os"
	"strconv"
)

func ReadBalance(money float64, b bool) {
	data, _ := os.ReadFile("balance")
	balance, _ := strconv.ParseFloat(string(data), 64)

	if b {
		WriteBalance(balance + money)
		return
	} else if money < balance {
		WriteBalance(balance - money)
		return
	}
	fmt.Println("столько на баласне нету")
}

func WriteBalance(balance float64) {
	file, _ := os.Create("balance")

	defer file.Close()

	fmt.Fprintf(file, "%.2f", balance)
}
