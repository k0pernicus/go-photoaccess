package utils

import (
	"net/http"

	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusNoContent,
		Response:   nil,
	})
}
