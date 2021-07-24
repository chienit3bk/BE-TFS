package middleware

import (
	"net/http"
	"project/packages/authentication/token"
	"project/packages/handlers/response"
)

func UserAuthorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//getToken
		tokenString, err := token.GetTokenString(r)
		if err != nil {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}

		//verify
		_, err2 := token.VerifyToken(tokenString)
		if err2 != nil {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err2.Error()})
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func UserAuthorize(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//getToken
// 		tokenString, err := token.GetTokenString(r)
// 		if err != nil {
// 			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
// 			return
// 		}

// 		//verify
// 		_, err2 := token.VerifyToken(tokenString)
// 		if err2 != nil {
// 			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err2.Error()})
// 			return
// 		}

// 		next(w, r)
// 	}
// }
