package photo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k0pernicus/go-photoaccess/internal/helpers"
	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

// Get returns the photo object, if it exists
func Get(w http.ResponseWriter, r *http.Request) {
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

	// TODO: return the object
	// _, exists := app.DB.Load(id)
	// if !exists {
	// 	fmt.Println("ID does not exists")
	// 	helpers.AnswerWith(w, types.ServiceResponse{
	// 		StatusCode: http.StatusNotFound,
	// 		Response: types.ExistsResponse{
	// 			Message: types.OK,
	// 		},
	// 	})
	// 	return
	// }

	helpers.AnswerWith(w, types.ServiceResponse{
		StatusCode: http.StatusOK,
		Response: types.GetResponse{
			Data: id,
		},
	})
}
