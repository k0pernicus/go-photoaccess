package photo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	app "github.com/k0pernicus/go-photoaccess/internal"
	"github.com/k0pernicus/go-photoaccess/internal/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Delete allows to delete a photo with associated metadata (or annotations)
func Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Cannot find 'id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	commandTag, err := app.DB.Exec(ctx, fmt.Sprintf("DELETE from %s where id=%s", photosTableName, id))
	if err != nil {
		log.Warningf("Warning when deleting the element with id %s on table %s: %+v", id, photosTableName, err)
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

	if err := db_ops.DeleteAllAnnotations(ctx, id); err != nil {
		log.Errorf("Error when trying to delete all annotations from photo %s", id)
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.DeleteResponse{
			Message: types.OK,
		},
	})
}
