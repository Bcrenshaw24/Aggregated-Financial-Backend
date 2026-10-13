package entities

import (
	"errors"
	"strings"
	"uuid"
)

// Recieving from API
type Transaction struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	UserID        uuid.UUID `json:"User_id"`
	Amount        float64   `json:"Amount"`
	Category      string    `json:"Category"`
	Timestamp     string    `json:"Timestamp"`
}

// Database Schema
type Account struct {
	UserID       uuid.UUID
	Email        string
	Password     string
	Balance      float64
	Transactions []*Transaction
	WeeklySpend  float64
	MonthlySpend float64
}

func (t Transaction) validate() error {
	if t.TransactionID == uuid.Nil() {
		return errors.New("Transaction ID missing")
	}
	if t.UserID == uuid.Max() {
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

func (a Account) validate() error {
	if a.UserID == uuid.Max() {
		return errors.New("User ID missing")
	}
	if strings.TrimSpace(a.Email) == "" {
		return errors.New("Email must be present")
	}
	if a.WeeklySpend < 0 {
		return errors.New("Cannot spend negative dollars")
	}
	if a.MonthlySpend < 0 {
		return errors.New("Cannot spend negative dollars")
	}
	return nil
}
