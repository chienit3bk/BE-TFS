package elastic_search

import (
	"fmt"
	"project/database"
)

func AutoSync() {
	//connect database
	db := *database.ConnectToDatabase()

	//Get the list of products that have not been saved on ES
	var listProducts []database.Product
	db.Where("ES = ?", false).Find(&listProducts)

	//connect to ES
	esClient, err := NewESClient(URL)
	if err != nil {
		fmt.Println("Cannot create new ESClient")
		return
	}
	productManager  := NewProductManager(esClient)

	for i, product := range listProducts{
		//save to ES
		productManager.AddProduct(&product)
		//mark synced
		listProducts[i].ES = true
	}
	//update database
	db.Save(&listProducts)
}
