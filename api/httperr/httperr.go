package httperr

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

func New(status int, message string, errors ...error) huma.StatusError {

	log.Error(message, "status", status, "err", errors)

	errId, err := uuid.NewUUID()
	if err != nil {
		return huma.NewError(http.StatusInternalServerError, "error proccessing error")
	}

	return huma.NewError(status, fmt.Sprintf("[%s] %s", errId, message))
}
