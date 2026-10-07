package main

import (
	"fmt"
	"net/http"
)

var Version = "0.0.2"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /transaction", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "All transactions")
	})

	mux.HandleFunc("GET /balance", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Total aggregated balance")
	})

	mux.HandleFunc("GET /spending/{timeframe}", func(w http.ResponseWriter, r *http.Request) {
		timeframe := r.PathValue("timeframe")
		fmt.Fprintf(w, "Spending over %s: 200", timeframe)
	})

	mux.HandleFunc("POST /transaction", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Push a new transaction")
	})

	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Listening on port:8000")
}
