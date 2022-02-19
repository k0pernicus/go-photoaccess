package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/k0pernicus/go-photoaccess/internal/annotation/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	photoID, ok := vars["photo_id"]
	if !ok {
		log.Debug("Cannot find 'photo_id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.MissingInformation,
			},
		})
		return
	}

	// Count the number of objects
	nbRows, err := db_ops.CountRows(ctx, photoID)
	if err != nil {
		log.Warningf("Warning when query count * for annotations: %+v", err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.ErrorResponse{
				Message: types.InternalError,
			},
		})
		return
	}

	if nbRows == 0 {
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusNotFound,
			Response: types.ErrorResponse{
				Message: types.EntityNotFound,
			},
		})
		return
	}

	annotations := make(types.Annotations, nbRows)
	if err := db_ops.GetAllAnnotations(ctx, &annotations, photoID); err != nil {
		log.Warningf("Warning when query * for annotations: %+v", err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.ErrorResponse{
				Message: types.InternalError,
			},
		})
		return
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.GetResponse{
			Data: annotations,
		},
	})
}

// Get returns the annotation object, if it exists
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	annotationID, ok := vars["annotation_id"]
	if !ok {
		log.Debug("Cannot find 'annotation_id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.MissingInformation,
			},
		})
		return
	}

	photoID, ok := vars["photo_id"]
	if !ok {
		log.Debug("Cannot find 'photo_id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.MissingInformation,
			},
		})
		return
	}

	var a types.Annotation
	err := db_ops.GetOneAnnotationWithKnownPhoto(r.Context(), annotationID, photoID, &a)

	if err == pgx.ErrNoRows {
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusNotFound,
			Response: types.ErrorResponse{
				Message: types.EntityNotFound,
			},
		})
		return
	}

	if err != nil {
		log.Warningf("Warning when query * for annotations with id %s: %+v", annotationID, err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.ErrorResponse{
				Message: types.InternalError,
			},
		})
		return
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.GetResponse{
			Data: a,
		},
	})
}
