package utils

import (
	"net/http"

	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func Version(w http.ResponseWriter, r *http.Request) {
	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.GetResponse{
			Data: app.Version,
		},
	})
}
