package user

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"project/database"
	userAuthen "project/packages/authentication/user"
	"project/packages/handlers/response"
	"project/packages/mail"
	"project/packages/redisCache"
	"strconv"
	"strings"
	"time"
)

type ForgotFormStep1 struct {
	Email string `json:"email"`
}
type ForgotFormStep2 struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}
type ForgotFormStep3 struct {
	Email   string `json:"email"`
	NewPass string `json:"newPass"`
}

//enter your email
func ForgotPasswordStep1(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newForgot ForgotFormStep1
	err := json.NewDecoder(r.Body).Decode(&newForgot)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//check email
	ok, message, userID, username := checkEmail(newForgot.Email)
	if !ok {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message})
		return
	}

	//create code
	code := createCode()

	//save code on rediscahe: key = userID, value = code
	rd := redisCache.Redis{}
	rd.Connect()
	rd.SetData(strconv.Itoa(userID), strconv.Itoa(code))

	//send code
	mail.SendCodeVerify(code, username, newForgot.Email)

	//response
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Your verrification code has been sent !"})
}

//enter code
func ForgotPasswordStep2(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newForgot ForgotFormStep2
	err := json.NewDecoder(r.Body).Decode(&newForgot)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//get userID
	_, _, userID, _ := checkEmail(newForgot.Email)

	//verify code
	rd := redisCache.Redis{}
	rd.Connect()
	codeSaved, err := rd.GetData(strconv.Itoa(userID))

	if err != nil || codeSaved != strconv.Itoa(newForgot.Code) {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Your code is incorrect !"})
		return
	}

	//delete code
	rd.DeleteData(strconv.Itoa(userID))

	//response
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "Your code is correct !"})
}

//enter new password
func ForgotPasswordStep3(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newForgot ForgotFormStep3
	err := json.NewDecoder(r.Body).Decode(&newForgot)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//check new pass
	ok, message := userAuthen.CheckPassword(newForgot.NewPass)
	if !ok {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": message})
		return
	}

	//get userID
	_, _, userID, _ := checkEmail(newForgot.Email)

	//connect database
	db := *database.ConnectToDatabase()

	//find by ID (Primary key)
	var user database.User
	db.First(&user, userID)

	//check user exist
	if user.ID == 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such user !"})
		return
	}

	//encrypt password
	passHashed, _ := userAuthen.Hash(newForgot.NewPass)

	//update and save
	user.Password = passHashed
	db.Save(&user)

	//reponse
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "User Password Updated !"})

}

//return isValid, message, idUser, username if exist
func checkEmail(email string) (bool, string, int, string) {
	if email != "" {
		//check form
		var arr = strings.Split(email, "@")
		if arr[len(arr)-1] != "gmail.com" {
			return false, "Invalid email", 0, ""
		}

		//check database
		db := database.ConnectToDatabase()
		var user database.User
		db.Where("email = ?", email).First(&user)
		if user.ID != 0 {
			//email already exists
			return true, "", user.ID, user.Name
		}

		return false, "This email is not registered", 0, ""
	}
	return false, "Email cannot be left blank", 0, ""
}

//Create verification code
func createCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	return rand.Intn(max-min+1) + min
}
