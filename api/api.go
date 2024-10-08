package api

import (
	"encoding/json"
	"net/http"

	"github.com/waksun0x00/todoAPI/internal/tools"
)

// type TodoParams struct {
// 	ID string
// }

type TodoResponse struct {
	// success code
	Code int

	Details tools.Todo
}

type TodoListResponse struct {
	Code int

	TodoList []tools.Todo
}

type Error struct {
	// error code
	Code int

	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
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
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
