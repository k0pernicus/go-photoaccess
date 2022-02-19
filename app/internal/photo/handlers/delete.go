package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	annotation_ops "github.com/k0pernicus/go-photoaccess/internal/annotation/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/internal/photo/db_ops"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Delete allows to delete a photo with associated metadata (or annotations)
func Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		log.Debugf("Cannot find 'id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	commandTag, err := db_ops.DeleteOne(ctx, id)
	if err != nil {
		log.Warningf("Warning when deleting the photo with id %s: %+v", id, err)
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

	if err := annotation_ops.DeleteAllAnnotations(ctx, id); err != nil {
		log.Errorf("Error when trying to delete all annotations from photo %s", id)
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusNoContent,
		Response:   nil,
	})
}
