package main

import (
	"fmt"
	"net/http"
)

type welcome string

func (wc welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Golang application!")
}

func main() {
	router := http.NewServeMux()

	var wc welcome

	router.Handle("/", wc)
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the login page")
	})

	server := http.Server{
		Addr:    ":8070",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server starting error: ", err)
	}
}