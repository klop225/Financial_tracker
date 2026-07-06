package main

import (
	"Financial_tracker/command"
	"fmt"
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
			if err := command.Add(money); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Средства успешно добавлены\n")
			}
		case "take":
			fmt.Print("сколько хотите снять? ")
			fmt.Scan(&money)
			if money <= 0 {
				fmt.Println("ошибка при снятие")
				continue
			}
			if err := command.Take(money); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Средства успешно сняты\n")
			}
		case "balance":
			balance, err := command.GetBalance()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("ваш баланс %.2f\n", balance)
			}
		case "history":
			history, err := command.GetHistory()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Print(history)
			}
		}
	}
}
