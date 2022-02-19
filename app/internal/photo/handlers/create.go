package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/internal/photo/db_ops"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Create permits to create a photo entity
func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var c types.PhotoCreationRequest
	err := decoder.Decode(&c)
	if err != nil {
		log.Debugf("Error when decoding message: %s", err.Error())
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.PostResponse{
				Data:    nil,
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	if strings.TrimSpace(c.Data) == "" {
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusUnprocessableEntity,
			Response: types.ErrorResponse{
				Message:   types.MissingInformation,
				ExtraInfo: "Empty data",
			},
		})
		return
	}

	id, err := db_ops.CreatePhoto(ctx, c)
	if err != nil {
		log.Errorf("cannot create Photo object due to error %+v", err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.PostResponse{
				Data:    nil,
				Message: types.InternalError,
			},
		})
		return
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.PostResponse{
			Data: types.PhotoCreationResponse{
				ID: strconv.Itoa(id),
			},
		},
	})
}
