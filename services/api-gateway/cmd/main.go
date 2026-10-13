package main

import (
	"finance/internal/entities"
	"finance/internal/services"
	"fmt"
	"net/http"
)

var Mock = true

var Version = "0.0.2"

func main() {
	mux := http.NewServeMux()
	users := services.UserService{
		Running:  true,
		Accounts: make(map[string]*entities.Account),
	}

	mux.HandleFunc("POST /account", users.HandleAccountCreation)

	mux.HandleFunc("GET /account/{email}", users.GetAccountDetails)
	fmt.Println("Listening on port:8000")
	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		fmt.Println(err.Error())
	}

}
