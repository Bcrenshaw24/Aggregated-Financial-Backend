package main

import (
	"finance/internal/entities"
	"finance/internal/services"
	"fmt"
	"net/http"
	"uuid"
)

var Mock = true

var Version = "0.0.2"

func main() {
	mux := http.NewServeMux()
	users := services.UserService{
		Running:  true,
		Accounts: make(map[uuid.UUID]*entities.Account),
	}

	mux.HandleFunc("POST /account", users.HandleAccountCreation)
	fmt.Println("Listening on port:8000")
	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		fmt.Println(err.Error())
	}

}
