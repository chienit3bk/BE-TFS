package product

import (
	"fmt"
	"net/http"
	"project/database"
	"project/packages/handlers/response"
)

//http://localhost:8080/products/resolution?value=8K
func FilterAdvanced(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok---------------------")
	//get value
	params := r.URL.Query()
	resolution := params["resolution"]
	technology := params["technology"]
	typee := params["type"]
	fmt.Println("technology:  ", technology)
	fmt.Println("resolution:  ", resolution)
	fmt.Println("type:  ", typee)

	//get main keyword
	var resolutionValue, technologyValue, typeValue string
	if len(resolution) != 0 {
		resolutionValue = resolution[0]
	} else {
		resolutionValue = "%%"
	}
	if len(technology) != 0 {
		technologyValue = technology[0]
	} else {
		technologyValue = "%%"
	}
	if len(typee) != 0 {
		typeValue = typee[0]
	} else {
		typeValue = "%%"
	}

	//connect database
	db := database.ConnectToDatabase()

	//find
	var ListProducts []database.Product
	db.Where("technology LIKE ? AND resolution LIKE ? AND type LIKE ?", technologyValue, resolutionValue, typeValue).Find(&ListProducts)

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
