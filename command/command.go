package command

import "Financial_tracker/storage"

func Add(money float64) {
	storage.WriteBalance(money)
}
