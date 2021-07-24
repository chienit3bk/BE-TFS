package order

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//get ID user
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["userId"])

	//connect database
	db := database.ConnectToDatabase()

	//find
	var listOrder []database.Order
	db.Where("id_user = ?", id).Find(&listOrder)

	//Add list order details
	for i, order := range listOrder {
		var listOrderDetails []database.OrderDetail
		db.Where("id_order = ?", order.ID).Find(&listOrderDetails)

		listOrder[i].ListOrderDetails = listOrderDetails
	}

	//response
	response.ResponseWithJson(w, http.StatusOK, listOrder)

}
