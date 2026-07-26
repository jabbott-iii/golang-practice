package internal

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// success code
	Code int

	// Balance
	Balance int64
}

type CreateUserParams struct {
	Username string
	AuthToken string
	Coins int
}

type CreateUserResponse struct {
	Code int

	Username string
	AuthToken string
	Coins int
}

type GetUserParams struct {
	Username string
}

type GetUserResponse struct {
	Code int

	Username string
	AuthToken string
	Coins int
}
type UpdateCoinParams struct {
	Username string
	Coins int
}

type UpdateCoinResponse struct {
	Code int

	Username string
	Coins int
}

// Error response
type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	var resp = Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)