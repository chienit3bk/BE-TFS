package order

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
	var newOrder database.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//get ID user
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["userId"])

	//connect database
	db := database.ConnectToDatabase()
	
	//get user's info
	var user database.User
	db.First(&user, id)

	//set
	newOrder.IDUser = id
	newOrder.CreatedAt = time.Now()
	newOrder.Username = user.Username
	newOrder.PhoneUser = user.Phone

	//save to database
	db.Create(&newOrder)

	//response
	response.ResponseWithJson(w, http.StatusCreated, newOrder)
}
