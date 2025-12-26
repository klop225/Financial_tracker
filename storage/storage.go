package storage

import (
	"fmt"
	"os"
)

func WriteBalance(money float64) {
	file, _ := os.Create("balance")

	defer file.Close()

	fmt.Fprintf(file, "%.2f", money)
}
