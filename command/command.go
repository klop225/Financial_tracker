package command

import "Financial_tracker/storage"

func Add(money float64) {
	storage.ReadBalance(money, true)
}

func Take(money float64) {
	storage.ReadBalance(money, false)
}
