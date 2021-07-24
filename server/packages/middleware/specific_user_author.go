package middleware

import (
	"net/http"

	"project/packages/authentication/token"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func SpecificUserAuthorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//getToken
		tokenString, err := token.GetTokenString(r)
		if err != nil {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}

		//verify
		verifiedtoken, err2 := token.VerifyToken(tokenString)
		if err2 != nil {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err2.Error()})
			return
		}

		//get token payload
		payloadData, err3 := token.ExtractTokenPayloadData(verifiedtoken)
		if err3 != nil {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err3.Error()})
			return
		}

		//get ID user
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["userId"])

		//check authorize
		if int(payloadData.UserId) != id {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "You are not authorized to do this !"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func SpecificUserAuthorize(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//getToken
// 		tokenString, err := token.GetTokenString(r)
// 		if err != nil {
// 			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
// 			return
// 		}

// 		//verify
// 		verifiedtoken, err2 := token.VerifyToken(tokenString)
// 		if err2 != nil {
// 			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err2.Error()})
// 			return
// 		}

// 		//get token payload
// 		payloadData, err3 := token.ExtractTokenPayloadData(verifiedtoken)
// 		if err3 != nil {
// 			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err3.Error()})
// 			return
// 		}

// 		//get ID user
// 		vars := mux.Vars(r)
// 		id, _ := strconv.Atoi(vars["id"])

// 		//check authorize
// 		if int(payloadData.UserId) != id  {
// 			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "You are not authorized to do this !"})
// 			return
// 		}

// 		next(w, r)
// 	}
// }
