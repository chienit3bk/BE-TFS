package product

import (
	"fmt"
	"net/http"
	// "project/database"
	"project/packages/handlers/response"
	es "project/packages/elastic_search"
)


func SearchByNameES(w http.ResponseWriter, r *http.Request) {
	es.AutoSync()
	//get name
	params := r.URL.Query()
	name := params["name"]
	

	//connect to ES
	esClient, err := es.NewESClient(es.URL)
	if err != nil {
		fmt.Println("Cannot create new ESClient")
		return
	}
	productManagerr  := es.NewProductManager(esClient)


	//find 
	listProducts := productManagerr.SearchProducts(name[0])

	//Check result
	if len(listProducts) == 0 {
		fmt.Println("Not Found !")
		response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Product Not Found!"})
		return
	}
	

	//response
	response.ResponseWithJson(w, http.StatusOK, listProducts)

	// fmt.Println(product)
}

