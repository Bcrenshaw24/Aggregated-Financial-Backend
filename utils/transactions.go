package utils

import (
	"errors"
	"strings"
)

type Transaction struct {
	Transaction_id string  `json:"transaction_id"`
	User_id        string  `json:"User_id"`
	Amount         float32 `json:"Amount"`
	Category       string  `json:"Category"`
	Timestamp      string  `json:"Timestamp"`
}

func (t Transaction) validate() error {
	if strings.TrimSpace(t.Transaction_id) == "" {
		return errors.New("Transaction ID missing")
	}
	if strings.TrimSpace(t.User_id) == "" {
		return errors.New("User ID missing")
	}
	if t.Amount < 0 {
		return errors.New("Amount cannot be less than 0")
	}
	if strings.TrimSpace(t.Category) == "" {
		return errors.New("Category cannot be none")
	}
	if strings.TrimSpace(t.Timestamp) == "" {
		return errors.New("Transaction must have a timestamp")
	}
	return nil
}
