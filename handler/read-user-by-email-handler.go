package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/api/adapter"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
)

type ReadUserByEmailHandler struct {
	userRepository UserRepository
}

func NewReadUserByEmailHandler(userRepository UserRepository) ReadUserByEmailHandler {
	return ReadUserByEmailHandler{
		userRepository,
	}
}

func (h ReadUserByEmailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	user := h.userRepository.FindUserByEmail(email)

	response, _ := json.Marshal(adapter.NewUserAdapter(user))
	w.Write(response)
}