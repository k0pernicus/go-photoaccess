package photo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/k0pernicus/go-photoaccess/internal/db_ops"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
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
		fmt.Printf("Error when decoding message: %s", err.Error())
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.PostResponse{
				Data:    nil,
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	if !c.IsValid() {
		log.Debugf("received invalid information from User %+v", err)
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.PostResponse{
				Data:    nil,
				Message: types.InvalidInformation,
			},
		})
		return
	}

	// Store the content in DB + retrieve the ID
	if c.IsAdditional && !db_ops.AnnotationExists(ctx, c.AnnotationID) {
		log.Debugf("invalid annotation")
		helpers.AnswerWith(w, types.ServiceResponse{
			StatusCode: http.StatusBadRequest,
			Response: types.PostResponse{
				Data:    nil,
				Message: types.InvalidInformation,
			},
		})
		return
	}

	if c.IsAdditional {
		log.Debugf("creating additional photo in DB with annotation ID %d...", c.AnnotationID)
	} else {
		log.Debugf("creating new photo in DB...")
	}

	id, err := db_ops.CreatePhoto(ctx, photosTableName, c)
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
