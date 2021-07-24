package order

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//get ID order
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["orderId"])

	//connect database
	db := *database.ConnectToDatabase()

	//find by ID (Primary key)
	var order database.Order
	db.First(&order, id)

	//check
	if order.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such order !"})
		return
	}

	//update and save
	order.State = "deleted"
	db.Save(&order)

	//reponse
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Order deleted !"})
}
