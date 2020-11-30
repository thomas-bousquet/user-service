package error_handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/errors"
	"net/http"
)

func WriteJSONErrorResponse(w http.ResponseWriter, err error, logger *logrus.Logger) {
	switch e := err.(type) {
	case ValidationError:
		logger.Error(e)
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err, logger)
	case UnexpectedError:
		logger.Error(e)
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err, logger)
	case AuthorizationError:
		logger.Error(e)
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err, logger)
	default:
		logger.Error(e)
		unexpectedError :=  NewUnexpectedError()
		w.WriteHeader(unexpectedError.HttpCode)
		doWriteError(w, unexpectedError, logger)
	}
}

func doWriteError(w http.ResponseWriter, error error, logger *logrus.Logger) {
	defaultErrorMessage := http.StatusText(http.StatusInternalServerError)
	defaultHTTPCode := http.StatusInternalServerError

	body, err := json.Marshal(error)

	if err != nil {
		logger.Error(err)
		http.Error(w, defaultErrorMessage, defaultHTTPCode)
	}

	_, writeError := w.Write(body)

	if writeError != nil {
		logger.Error(writeError)
		http.Error(w, defaultErrorMessage, defaultHTTPCode)
	}
}
