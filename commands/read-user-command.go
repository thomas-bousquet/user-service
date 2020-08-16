package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/api/adapters"
	. "github.com/thomas-bousquet/startup/repositories"
	"net/http"
)

type ReadUserCommand struct {
	userRepository UserRepository
}

func NewReadUserCommand(userRepository UserRepository) ReadUserCommand {
	return ReadUserCommand{
		userRepository,
	}
}

func (c ReadUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := c.userRepository.FindUser(id)

	if err != nil {
		return err
	}

	response, err := json.Marshal(adapters.NewUserAdapter(user))

	if err != nil {
		return err
	}
	_, err = w.Write(response)

	return err
}