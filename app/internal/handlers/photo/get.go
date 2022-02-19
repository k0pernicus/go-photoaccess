package photo

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/k0pernicus/go-photoaccess/internal/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Get returns the photo object, if it exists
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		log.Debug("Cannot find 'id' query parameter in user's request")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	var p types.Photo
	err := db_ops.GetOnePhoto(r.Context(), photosTableName, id, &p)

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
		log.Warningf("Warning when query * on table %s with id %s: %+v", photosTableName, id, err)
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
			Data: p,
		},
	})
}
