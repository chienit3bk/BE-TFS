package user

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func GetInfoByID(w http.ResponseWriter, r *http.Request) {
	//get ID order
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["userId"])

	//connect database
	db := database.ConnectToDatabase()

	//find by ID (Primary key)
	var user database.User
	db.First(&user, id)

	//Check result
	if user.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "User Not Exist!"})
		return
	}

	//response
	response.ResponseWithJson(w, http.StatusOK, user)

}
