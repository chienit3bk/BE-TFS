package response

import (
	"encoding/json"
	"net/http"
)

//response with http Status Code and Message or Object
func ResponseWithJson(w http.ResponseWriter, httpStatusCode int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(httpStatusCode)
	//response with json
	json.NewEncoder(w).Encode(object)
}
