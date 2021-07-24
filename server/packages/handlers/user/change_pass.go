package user

import (
	"encoding/json"
	"net/http"
	"project/database"
	userAuthen "project/packages/authentication/user"
	"project/packages/handlers/response"
)

type PasswordUpdate struct {
	IDUser  int    `json:"id_user"`
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newUpdate PasswordUpdate
	err := json.NewDecoder(r.Body).Decode(&newUpdate)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := *database.ConnectToDatabase()

	//find by ID (Primary key)
	var user database.User
	db.First(&user, newUpdate.IDUser)

	//check user exist
	if user.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such user !"})
		return
	}

	//compare newpass and oldpass
	if newUpdate.NewPass == newUpdate.OldPass {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "NewPass must be different from Oldpass !"})
		return
	}

	//check old pass
	err2 := userAuthen.CheckPasswordHash(user.Password, newUpdate.OldPass)
	if err2 != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "OldPass Incorrect !"})
		return
	}

	//check new pass
	ok, message := userAuthen.CheckPassword(newUpdate.NewPass)
	if !ok {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message})
		return
	}

	//update and save
	passHashed, _ := userAuthen.Hash(newUpdate.NewPass)
	user.Password = passHashed

	db.Save(&user)

	//reponse
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "User Password Updated !"})

}
