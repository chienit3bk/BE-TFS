package user

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/authentication/token"
	"project/packages/handlers/response"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newLogin LoginForm
	err := json.NewDecoder(r.Body).Decode(&newLogin)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//check authentication
	ok, message, userId, passwordInDatabase, isAdmin := checkUsernameLogin(newLogin.Username)
	if ok {
		if ok2, message2 := checkPasswordLogin(newLogin.Password, passwordInDatabase); !ok2 {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message2})
			return
		}
	} else {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message})
		return
	}

	//create token
	token, _ := token.CreateToken(userId, newLogin.Username, isAdmin)

	//response
	response.ResponseWithJson(w, http.StatusOK, token)
}

//return isValid, message, userId, passwordInDatabase and isAdmin if username already exist
func checkUsernameLogin(username string) (bool, string, int, string, bool) {
	if username != "" {
		//check database
		db := database.ConnectToDatabase()
		var user database.User
		db.Where("username = ?", username).First(&user)

		if user.ID != 0 {
			return true, "", user.ID, user.Password, user.IsAdmin
		}
		//if username not found
		return false, "Username or password incorrect", 0, "", false
	}

	return false, "Username cannot be left blank", 0, "", false
}

// return isValid and message
func checkPasswordLogin(password, hashedPassword string) (bool, string) {
	if password != "" {
		if len(password) < 8 {
			return false, "Password must be at least 8 characters"
		}
		//check
		err := CheckPasswordHash(hashedPassword, password)
		if err != nil {
			return false, "Username or password incorrect"
		}
		return true, ""
	}

	return false, "Password cannot be left blank"
}
