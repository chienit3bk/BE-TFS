package user

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strings"
)

type RegisterForm struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newRegister RegisterForm
	err := json.NewDecoder(r.Body).Decode(&newRegister)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//check email
	if ok, message := checkEmail(newRegister.Email); ok {
		//check username
		if ok2, message2 := checkUsername(newRegister.Username); ok2 {
			//check password
			if ok3, message3 := CheckPassword(newRegister.Password); !ok3 {
				response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message3})
				return
			}
		} else {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message2})
			return
		}
	} else {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message})
		return
	}

	//connect database
	db := database.ConnectToDatabase()

	//encrypt password
	password, _ := Hash(newRegister.Password)

	//Create new user
	var newUser = database.User{
		Username: newRegister.Username,
		Password: password,
		Email:    newRegister.Email,
		Phone:    newRegister.Phone,
		IsAdmin:  false,
	}

	//save to database
	db.Create(&newUser)

	//response
	response.ResponseWithJson(w, http.StatusCreated, map[string]string{"message": "Register Successfully !"})

}

// return isValid and message
func checkEmail(email string) (bool, string) {
	if email != "" {
		//check form
		var arr = strings.Split(email, "@")
		if arr[len(arr)-1] != "gmail.com" {
			return false, "Invalid email"
		}
		//check database
		db := database.ConnectToDatabase()
		var user database.User
		db.Where("email = ?", email).First(&user)
		if user.ID != 0 {
			//email already exists
			return false, "This email has been used to register for another account"
		}

		return true, "nothing"
	}
	return false, "Email cannot be left blank"
}

// return isValid and message
func checkUsername(username string) (bool, string) {
	if username != "" {
		//check database
		db := database.ConnectToDatabase()
		var user database.User
		db.Where("username = ?", username).First(&user)
		if user.ID != 0 {
			//username already exists
			return false, "Username already exists"
		}
		return true, "nothing"
	}
	return false, "Username cannot be left blank"
}

// return isValid and message
func CheckPassword(password string) (bool, string) {
	if password != "" {
		if len(password) < 8 {
			return false, "Password must be at least 8 characters"
		}
		return true, "nothing"
	}
	return false, "Password cannot be left blank"
}
