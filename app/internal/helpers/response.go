package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/k0pernicus/go-photoaccess/pkg/types"
)

func AnswerWith(w http.ResponseWriter, response types.ServiceResponse) {
	w.WriteHeader(response.StatusCode)
	b, _ := json.Marshal(response)
	w.Write(b)
}

func AnswerRaw(w http.ResponseWriter, response types.ServiceResponse) {
	w.WriteHeader(response.StatusCode)
	w.Write([]byte(response.Response.(string)))
}
