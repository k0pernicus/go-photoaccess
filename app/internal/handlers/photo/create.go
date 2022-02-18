package photo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Create permits to create a photo entity
func Create(w http.ResponseWriter, r *http.Request) {
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

	// Store the content in DB + retrieve the ID
	if c.IsAdditional {
		log.Debug("creating additional photo...")
		// TODO: Check for the annotation entity
	}

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.PostResponse{
			Data: types.PhotoCreationResponse{
				ID: "xxx", // TODO: change for returned ID from DB
			},
		},
	})
}
