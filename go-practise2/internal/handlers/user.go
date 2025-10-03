package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type GetUserResponse struct {
	UserID int `json:"user_id"`
}

type CreateUserRequest struct {
	Name string `json:"name" `
}

type CreateUserResponse struct {
	Created string `json:"created"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	//get and validate id
	idStr := r.URL.Query().Get("id")
	if idStr == ""{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid id"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid id"})
		return
	}

	//success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetUserResponse{UserID: id})
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	//parse json body
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid name"})
		return
	}

	//success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{Created: req.Name})
}