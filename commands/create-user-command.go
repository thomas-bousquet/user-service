package commands

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/models"
	. "github.com/thomas-bousquet/user-service/repositories"
	"github.com/thomas-bousquet/user-service/utils/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type CreateUserCommand struct {
	userRepository UserRepository
	validator      validator.Validator
}

func NewCreateUserCommand(userRepository UserRepository, validator validator.Validator) CreateUserCommand {
	return CreateUserCommand{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (c CreateUserCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) *errors.Error {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		logger.Error(err)
		return errors.NewUnexpectedError()
	}

	validationErrors := c.validator.ValidateStruct(user)

	if len(validationErrors) > 0 {
		return errors.NewValidationError("An error occurred when validating user fields", validationErrors)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Errorf("error when encrypting new user's password with email %q: %v", user.Email, err)
		return errors.NewUnexpectedError()
	}

	user.Password = string(encryptedPassword)
	userId, err := c.userRepository.CreateUser(user)

	if err != nil {
		logger.Errorf("error creating new user: %v", err)
		return errors.NewUnexpectedError()
	}

	response, err := json.Marshal(map[string]primitive.ObjectID{"id": userId})

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		return errors.NewUnexpectedError()
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)

	if err != nil {
		logger.Errorf("error writing response: %v", err)
		return errors.NewUnexpectedError()
	}

	return nil
}
