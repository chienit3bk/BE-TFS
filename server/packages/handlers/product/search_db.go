package product

import (
	"fmt"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
)

func SearchByNameDB(w http.ResponseWriter, r *http.Request) {
	//get name
	params := r.URL.Query()
	name := params["name"]

	//connect database
	db := database.ConnectToDatabase()

	//find
	var listProducts []database.Product
	namesearch := "%" + name[0] + "%"
	db.Where("name like ?", namesearch).Find(&listProducts)

	//Check result
	if len(listProducts) == 0 {
		fmt.Println("Not Found !")
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Product Not Found!"})
		return
	}
	for i, product := range listProducts {
		// var listDescriptions []database.Description
		var listImages []database.Image
		var listOptions []database.Option

		// db.Where("id_product = ?", product.ID).Find(&listDescriptions)
		db.Where("id_product = ?", product.ID).Find(&listImages)
		db.Where("id_product = ?", product.ID).Find(&listOptions)

		//bắt buộc phải dùng ListProducts[i]. thay cho product. bởi vì lúc đấy mới cho sửa slice, tương tự cho vòng for dòng 35
		// listProducts[i].ListDescriptions = listDescriptions
		listProducts[i].ListImages = listImages
		listProducts[i].ListOptions = listOptions
		// for j, op := range listProducts[i].ListOptions {
		// 	var variant database.Variant
		// 	db.Where("id_option = ?", op.ID).Find(&variant)
		// 	listProducts[i].ListOptions[j].Variant = variant
		// }
	}

	//response
	response.ResponseWithJson(w, http.StatusOK, listProducts)

	// fmt.Println(product)
}
