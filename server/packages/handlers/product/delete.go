package product

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//get id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["productId"])
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid ID !"})
		return
	}

	//connect database
	db := database.ConnectToDatabase()

	//Update quantity to 0
	db.Model(&database.Option{}).Where("id_product = ?", id).Update("quantity", 0)

	//response
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Delete Product Successful !"})
}
