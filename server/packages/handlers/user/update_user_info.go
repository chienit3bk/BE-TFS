package user

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var updateUser database.User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	//get ID user
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["userId"])

	//connect database
	db := *database.ConnectToDatabase()

	//find by ID (Primary key)
	var user database.User
	db.First(&user, id)

	//check
	if user.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such user !"})
		return
	}

	//safe
	isAdmin := user.IsAdmin
	password := user.Password
	username := user.Username

	//update and save
	user = updateUser
	user.ID = id
	user.Username = username
	user.IsAdmin = isAdmin
	user.Password = password
	db.Save(&user)

	//reponse
	response.ResponseWithJson(w, http.StatusOK, user)

}
