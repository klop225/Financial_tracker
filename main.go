package main

import (
	"Financial_tracker/command"
	"fmt"
	"os"
)

func main() {
	var cmd string
	var money float64

	for {
		fmt.Scan(&cmd)
		switch cmd {
		case "exit":
			fmt.Println("bye")
			return
		case "add":
			fmt.Print("какая сумма пополнения ")
			fmt.Scan(&money)
			if money <= 0 {
				fmt.Println("ошибка при пополнение")
				continue
			}
			command.Add(money)
		case "balance":
			data, _ := os.ReadFile("balance")
			fmt.Printf("ваш баланс %s", string(data))
		}
	}
}
