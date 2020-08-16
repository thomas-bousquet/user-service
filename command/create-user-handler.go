package command

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/custom-error"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"github.com/thomas-bousquet/startup/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CreateUserCommand struct {
	userRepository UserRepository
	validator      utils.Validator
}

func NewCreateUserCommand(userRepository UserRepository, validator utils.Validator) CreateUserCommand {
	return CreateUserCommand{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (h CreateUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	var user = User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	errors := h.validator.ValidateStruct(user)

	if len(errors) > 0 {
		return NewValidationError(errors)
	}

	userId, err := h.userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	response, err := json.Marshal(map[string]primitive.ObjectID{"id": userId})

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	 _, err = w.Write(response)

	return err
}
