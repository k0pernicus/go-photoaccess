package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k0pernicus/go-photoaccess/internal/annotation/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Delete allows to delete a photo with associated metadata (or annotations)
func Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	annotationID, ok := vars["annotation_id"]
	if !ok {
		log.Debug("Cannot find 'annotation_id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ErrorResponse{
				Message:   types.MissingInformation,
				ExtraInfo: "missing 'annotation_id' query parameter",
			},
		})
		return
	}

	photoID, ok := vars["photo_id"]
	if !ok {
		log.Debug("Cannot find 'photo_id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ErrorResponse{
				Message:   types.MissingInformation,
				ExtraInfo: "missing 'photo_id' query parameter",
			},
		})
		return
	}

	commandTag, err := db_ops.DeleteOne(ctx, annotationID, photoID)
	if err != nil {
		log.Warningf("Warning when deleting the element with id %s for annotations: %+v", annotationID, err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.ErrorResponse{
				Message: types.InternalError,
			},
		})
		return
	}

	// No rows have been affected (not found)
	if commandTag.RowsAffected() != 1 {
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusNotFound,
			Response: types.ErrorResponse{
				Message: types.EntityNotFound,
			},
		})
		return
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusNoContent,
		Response:   nil,
	})

}
