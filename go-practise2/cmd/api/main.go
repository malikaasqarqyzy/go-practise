package main

import (
	"log"
	"net/http"

	"go-practise2/internal/handlers"
	"go-practise2/internal/middleware"
)

func main(){
	mux := http.NewServeMux()

	//Register routes
	mux.HandleFunc("GET /user", handlers.GetUser)
	mux.HandleFunc("POST /user", handlers.CreateUser)

	//Apply middleware chain
	handler := middleware.AuthMiddleware(mux)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}