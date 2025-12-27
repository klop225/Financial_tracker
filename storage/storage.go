package storage

import (
	"fmt"
	"os"
	"strconv"
)

func ReadBalance(money float64) {
	data, _ := os.ReadFile("balance")
	balance, _ := strconv.ParseFloat(string(data), 64)

	WriteBalance(balance + money)
}

func WriteBalance(balance float64) {
	file, _ := os.Create("balance")

	defer file.Close()

	fmt.Fprintf(file, "%.2f", balance)
}
