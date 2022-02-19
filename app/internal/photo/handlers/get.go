package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	annotation_db_ops "github.com/k0pernicus/go-photoaccess/internal/annotation/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/internal/photo/db_ops"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// The key query parameter to include all annotations in the response
const includeAnnotationsQueryName string = "include_annotations"

func GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Count the number of objects
	nbRows, err := db_ops.CountRows(ctx)
	if err != nil {
		log.Warningf("Warning when query count * for photos: %+v", err)
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

	photos := make([]types.Photo, nbRows)
	if err := db_ops.GetAllPhotos(ctx, &photos); err != nil {
		log.Warningf("Warning when query * for photos: %+v", err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.ErrorResponse{
				Message: types.InternalError,
			},
		})
		return
	}

	includeAnnotations := r.URL.Query().Get(includeAnnotationsQueryName)
	if strings.ToLower(includeAnnotations) == "true" {
		for index, photo := range photos {
			if nbRows, _ := annotation_db_ops.CountRows(ctx, strconv.Itoa(photo.ID)); nbRows > 0 {
				annotations := make(types.Annotations, nbRows)
				if err := annotation_db_ops.GetAllAnnotations(ctx, &annotations, strconv.Itoa(photo.ID)); err != nil {
					log.Warningf("cannot retrieve an annotation due with photo ID %d to error: %+v", photo.ID, err)
					continue
				}
				photos[index].Annotations = annotations
			}
		}
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.GetResponse{
			Data: photos,
		},
	})
}

// Get returns the photo object, if it exists
func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
	err := db_ops.GetOnePhoto(r.Context(), id, &p)

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
		log.Warningf("Warning when query * for photos with id %s: %+v", id, err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusInternalServerError,
			Response: types.ErrorResponse{
				Message: types.InternalError,
			},
		})
		return
	}

	includeAnnotations := r.URL.Query().Get(includeAnnotationsQueryName)
	if strings.ToLower(includeAnnotations) == "true" {
		if nbRows, _ := annotation_db_ops.CountRows(ctx, strconv.Itoa(p.ID)); nbRows > 0 {
			annotations := make(types.Annotations, nbRows)
			if err := annotation_db_ops.GetAllAnnotations(ctx, &annotations, strconv.Itoa(p.ID)); err == nil {
				log.Warningf("cannot retrieve an annotation due with photo ID %d to error: %+v", p.ID, err)
			} else {
				p.Annotations = annotations
			}
		}

	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.GetResponse{
			Data: p,
		},
	})
}
