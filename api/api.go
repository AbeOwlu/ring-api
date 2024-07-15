package api

import (
	"encoding/json"
	"net/http"
)

// query param by username
type CoinBalanceParam struct {
	Username string
}

// response and coin balance for user
type CoinBalanceResponse struct {
	Balance int64

	RespondCode int
}

// error message if invalid request
type ErrorRes struct {
	ResponseCode int

	ErrMessage string
}

func writeError(w http.ResponseWriter, messge string, code int) {
	resp := ErrorRes{
		ResponseCode: code,
		ErrMessage:   messge,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "Error processing Request", http.StatusInternalServerError)
	}
)
