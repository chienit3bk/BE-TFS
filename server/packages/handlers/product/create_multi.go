package product

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"project/packages/handleFile"
	"project/packages/handlers/response"
	"project/packages/realtime"
	"time"
)

type FileString struct {
	File string `json:"file"`
}

// chưa thực hiện lấy được file ở client gửi về, chỉ mới  chạy được ở file local
func CreateMulti(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var file FileString
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid bodyyyy"})
		return
	}

	//random file Name
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	a := rand.Intn(max-min+1) + min
	fileName := fmt.Sprintf("tmp/a%v.txt", a)

	//save savev
	handleFile.WriteToFile(fileName, file.File)

	// Báo có file mới
	realtime.NewFileName <- fileName

	//code
	response.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "File is being processed"})

}
