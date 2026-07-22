package main

import (
	"Financial_tracker/command"
	"Financial_tracker/storage"
	"fmt"
)

func main() {
	store := storage.NewFile("db/balance.json", "db/history.json")

	cmdHandler := command.NewCommandHandler(store)

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
			if err := cmdHandler.Add(money); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Средства успешно добавлены\n")
			}
		case "take":
			fmt.Print("сколько хотите снять? ")
			fmt.Scan(&money)
			if err := cmdHandler.Take(money); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Средства успешно сняты\n")
			}
		case "balance":
			balance, err := cmdHandler.GetBalance()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("ваш баланс %.2f\n", balance)
			}
		case "history":
			history, err := cmdHandler.GetHistory()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Print(history)
			}
		}
	}
}
