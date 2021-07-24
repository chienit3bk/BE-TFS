package product

import (
	"fmt"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//connect database
	db := database.ConnectToDatabase()

	//find
	var ListProducts []database.Product
	db.Find(&ListProducts)

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
		// for j, op := range ListProducts[i].ListOptions {
		// 	var variant database.Variant
		// 	db.Where("id_option = ?", op.ID).Find(&variant)
		// 	ListProducts[i].ListOptions[j].Variant = variant
		// }
	}

	//response
	response.ResponseWithJson(w, http.StatusOK, ListProducts)

	// fmt.Println(ListProducts)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	//get ID
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//connect database
	db := database.ConnectToDatabase()

	//find by ID (Primary key)
	var product database.Product
	db.First(&product, id)

	//Check result
	if product.ID == 0 {
		fmt.Println("Not Found !")
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Product Not Found!"})
		return
	}
	//Add details
	var listDescriptions []database.Description
	var listImages []database.Image
	var listOptions []database.Option

	db.Where("id_product = ?", product.ID).Find(&listDescriptions)
	db.Where("id_product = ?", product.ID).Find(&listImages)
	db.Where("id_product = ?", product.ID).Find(&listOptions)

	product.ListDescriptions = listDescriptions
	product.ListImages = listImages
	product.ListOptions = listOptions

	for j, op := range product.ListOptions {
		var variant database.Variant
		db.Where("id_option = ?", op.ID).Find(&variant)
		// variant./Order_Detail =
		product.ListOptions[j].Variant = variant
	}

	//response
	response.ResponseWithJson(w, http.StatusOK, product)

	// fmt.Println(product)
}
