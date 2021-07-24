package product

import (
	"fmt"
	"net/http"
	"project/database"
	"project/packages/handlers/response"

	"github.com/gorilla/mux"
)

//http://localhost:8080/products/resolution?value=8K
func Filter(w http.ResponseWriter, r *http.Request) {
	//get category
	vars := mux.Vars(r)
	categogy := vars["category"]
	fmt.Println(categogy)

	//get value
	params := r.URL.Query()
	value := params["value"]

	//connect database
	db := database.ConnectToDatabase()

	//find
	var ListProducts []database.Product
	db.Where(fmt.Sprintf("%v = ?", categogy), value[0]).Find(&ListProducts)

	//Add details
	for i, product := range ListProducts {
		// var listDescriptions []database.Description
		var listImages []database.Image
		var listOptions []database.Option

		// db.Where("id_product = ?", product.ID).Find(&listDescriptions)
		db.Where("id_product = ?", product.ID).Find(&listImages)
		db.Where("id_product = ?", product.ID).Find(&listOptions)

		//bắt buộc phải dùng ListProducts[i]. thay cho product. bởi vì lúc đấy mới cho sửa slice, tương tự cho vòng for dòng 35
		// ListProducts[i].ListDescriptions = listDescriptions
		ListProducts[i].ListImages = listImages
		ListProducts[i].ListOptions = listOptions
		for j, op := range ListProducts[i].ListOptions {
			var variant database.Variant
			db.Where("id_option = ?", op.ID).Find(&variant)
			ListProducts[i].ListOptions[j].Variant = variant
		}
	}

	//response
	response.ResponseWithJson(w, http.StatusOK, ListProducts)

}
