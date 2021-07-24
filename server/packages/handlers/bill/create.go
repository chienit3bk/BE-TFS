package bill

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var orderInfo database.Order
	err := json.NewDecoder(r.Body).Decode(&orderInfo)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//get ID user
	vars := mux.Vars(r)
	idUser, _ := strconv.Atoi(vars["userId"])

	//connect database
	db := database.ConnectToDatabase()

	//get full order's info
	var order database.Order
	db.First(&order, orderInfo.ID)

	//create bill
	var newBill = database.Bill{
		IDOrder:     orderInfo.ID,
		IDUser:      idUser,
		IdAdmin:     1,
		Username:    order.Username,
		AdminName:   "ngoc",
		PhoneUser:   order.PhoneUser,
		TotalPrice:  order.TotalPrice,
		AddressUser: order.AddressUser,
		State:       "unpaid",
		CreatedAt:   time.Now(),
	}

	//save to database
	db.Create(&newBill)

	//getBill
	var bill database.Bill
	db.Where("id_order = ?", newBill.IDOrder, "create_at = ?", newBill.CreatedAt).Last(&bill)

	//response
	response.ResponseWithJson(w, http.StatusCreated, bill)
}
