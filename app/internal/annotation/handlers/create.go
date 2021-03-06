package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/k0pernicus/go-photoaccess/internal/annotation/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	photo_ops "github.com/k0pernicus/go-photoaccess/internal/photo/db_ops"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Create permits to create an annotation entity
func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var a types.AnnotationCreationRequest
	err := decoder.Decode(&a)

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

	// Get from the router the photo_id information
	vars := mux.Vars(r)
	photoID, ok := vars["photo_id"]

	if a.Text == "" || !ok || !photo_ops.PhotoExists(ctx, photoID) {
		log.Debugf("invalid information (with photo id %s)", photoID)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusUnprocessableEntity,
			Response: types.ErrorResponse{
				Message:   types.InvalidInformation,
				ExtraInfo: "The data are invalid (missing annotation text, or associated photo does not exists)",
			},
		})
		return
	}

	photoIDNum, _ := strconv.Atoi(photoID)
	id, err := db_ops.CreateAnnotation(ctx, a, photoIDNum)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			// For the "unique" constraint
			case "23505":
				helpers.AnswerWith(w, types.ServiceResponse{
					StatusCode: http.StatusConflict,
					Response: types.PostResponse{
						Data:    nil,
						Message: types.AlreadyExists,
					},
				})
				return
			}
		}
		// else
		log.Errorf("cannot create Annotation object due to error %+v", err)
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
			Data: types.AnnotationCreationResponse{
				ID: strconv.Itoa(id),
			},
		},
	})
}
