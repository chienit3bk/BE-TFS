package user

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
)

//for admin
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	//connect database
	db := database.ConnectToDatabase()

	//find
	var listUsers []database.User
	db.Find(&listUsers)

	//response
	response.ResponseWithJson(w, http.StatusOK, listUsers)
}
