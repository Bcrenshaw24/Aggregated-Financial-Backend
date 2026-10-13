package services

import (
	"encoding/json"
	"finance/internal/entities"
	"fmt"
	"net/http"
	"sync"
	"uuid"
)

type UserService struct {
	Running  bool
	mu       sync.RWMutex
	Accounts map[string]*entities.Account
}

type CreateAccountRequest struct {
	Balance  float64 `json:"balance"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func (u *UserService) HandleAccountCreation(w http.ResponseWriter, r *http.Request) {
	u.mu.Lock()
	defer u.mu.Unlock()
	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userID := uuid.New()
	email := req.Email
	pass := req.Password
	balance := req.Balance

	account := entities.Account{
		UserID:       userID,
		Email:        email,
		Password:     pass,
		Balance:      balance,
		Transactions: nil,
		WeeklySpend:  0.0,
		MonthlySpend: 0.0,
	}
	if _, ok := u.Accounts[email]; ok {
		fmt.Fprintf(w, "Account Already Exists!")
		return
	}
	u.Accounts[email] = &account
	fmt.Fprintf(w, "Account Created Successfully!")
}

func (u *UserService) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	if _, ok := u.Accounts[email]; ok {
		response := map[string]any{
			"status":        "success",
			"user_id":       u.Accounts[email].UserID,
			"email":         u.Accounts[email].Email,
			"password":      u.Accounts[email].Password,
			"balance":       u.Accounts[email].Balance,
			"weekly_spend":  u.Accounts[email].WeeklySpend,
			"monthly_spend": u.Accounts[email].MonthlySpend,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	fmt.Fprintf(w, "Account Not Found!")
}
