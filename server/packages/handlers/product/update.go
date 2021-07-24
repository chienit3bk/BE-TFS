package product

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var updateProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := *database.ConnectToDatabase()

	//find by ID (Primary key)
	var product database.Product
	db.First(&product, updateProduct.ID)

	//check
	if product.ID == 0 {
		// fmt.Println("There is no such product !")
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "There is no such product !"})
		return
	}

	//update and save
	product = updateProduct
	product.ES = true
	db.Save(&product)

	// //sync ES
	// es.AutoSync()

	//reponse
	response.ResponseWithJson(w, http.StatusOK, product)

}
