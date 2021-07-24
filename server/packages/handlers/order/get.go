package order

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	//get ID order
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["orderId"])

	//connect database
	db := database.ConnectToDatabase()

	//find by ID (Primary key)
	var order database.Order
	db.First(&order, id)

	//Check result
	if order.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Order Not Found!"})
		return
	}

	//Add list order details
	var listOrderDetails []database.OrderDetail
	db.Where("id_order = ?", order.ID).Find(&listOrderDetails)
	order.ListOrderDetails = listOrderDetails

	//response
	response.ResponseWithJson(w, http.StatusOK, order)

}
