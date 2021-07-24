package token

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	Authorized bool
	UserId     uint64
	Username   string
	IsAdmin    bool
	Exp        uint64
}

func CreateToken(userId int, username string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["username"] = username
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 15).Unix() //Token hết hạn sau 15 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func GetTokenString(r *http.Request) (string, error) {
	token := r.Header.Get("token")
	if token != "" {
		return token, nil
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", fmt.Errorf("token not found")
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	//check valid
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return token, nil
}

//get payload data
func ExtractTokenPayloadData(verifiedToken *jwt.Token) (*TokenPayload, error) {
	claims, _ := verifiedToken.Claims.(jwt.MapClaims)

	authorized, ok := claims["authorized"].(bool) //boolean type
	if !ok {
		return nil, fmt.Errorf("token don't contain 'authorized'")
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userId"]), 10, 64) //int type
	if err != nil {
		return nil, fmt.Errorf("token don't contain 'userId'")
	}

	username, ok1 := claims["username"].(string) //string type
	if !ok1 {
		return nil, fmt.Errorf("token don't contain 'username'")
	}

	isAdmin, ok2 := claims["isAdmin"].(bool)
	if !ok2 {
		return nil, fmt.Errorf("token don't contain 'isAdmin'")
	}

	exp, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("token don't contain 'exp'")
	}

	return &TokenPayload{
		Authorized: authorized,
		UserId:     userId,
		Username:   username,
		IsAdmin:    isAdmin,
		Exp:        exp,
	}, nil
}

//get isAdmin's value
func ExtractTokenRole(verifiedToken *jwt.Token) (bool, error) {
	tokenPayload, err := ExtractTokenPayloadData(verifiedToken)
	if err != nil {
		return false, err
	}
	return tokenPayload.IsAdmin, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
