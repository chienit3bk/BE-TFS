package bill

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
)

func Cancel(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var updateBill database.Bill
	err := json.NewDecoder(r.Body).Decode(&updateBill)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := *database.ConnectToDatabase()

	//find by ID (Primary key)
	var bill database.Bill
	db.First(&bill, updateBill.ID)

	//check
	if bill.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such bill !"})
		return
	}

	//update state
	bill.State = "cancelled"

	//save
	db.Save(&bill)

	//reponse
	response.ResponseWithJson(w, http.StatusOK, bill)

}
