package product

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	es "project/packages/elastic_search"
)

func Create(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	var newProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	//connect database
	db := database.ConnectToDatabase()

	//check duplicate
	// Get first matched record
	var product database.Product
	db.Where("name = ?", newProduct.Name).First(&product)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	if product.ID != 0 {
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Product's name is Exist !"})
		return
	}

	//save product to database
	db.Create(&newProduct)

	//response
	response.ResponseWithJson(w, http.StatusCreated, newProduct)

	//create variants
	//get productId
	var product2 database.Product
	db.Where("name = ?", newProduct.Name).Last(&product2)
	//create variants from listOptions
	for _, op := range newProduct.ListOptions {
		//get optionId
		var option database.Option
		db.Where("id_product = ?", product2.ID, "price = ?", op.Price).First(&option)
		var variant = database.Variant{
			IDProduct:   product2.ID, //option.IDProduct
			IDOption:    option.ID,
			ProductName: product2.Name,
			Size:        option.Size,
			Price:       option.Price,
			SalePrice:   option.SalePrice,
		}
		db.Create(&variant)
	}

	//sync ES 
	es.AutoSync()
}
