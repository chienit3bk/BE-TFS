package elastic_search

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"project/database"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type ProductManager struct {
	esClient *ESClient
}

func NewProductManager(es *ESClient) *ProductManager {
	return &ProductManager{esClient: es}
}

func (pm *ProductManager) SearchProducts(name string) []database.Product {
	ctx := context.Background()

	//check
	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return nil
	}

	// build query to search for title
	query := elastic.NewSearchSource()
	query.Query(elastic.NewMatchQuery("name", name))

	// get search's service
	searchService := pm.esClient.
		Search().
		Index(indexName).
		SearchSource(query)

	// perform search query
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Cannot perform search with ES", err)
		return nil
	}

	// get result
	var listProducts []database.Product
	for _, hit := range searchResult.Hits.Hits {
		var product database.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			fmt.Println("Get data error: ", err)
			continue
		}
		fmt.Println(product)
		listProducts = append(listProducts, product)
	}

	return listProducts
}

func (pm *ProductManager) AddProduct(newProduct *database.Product) error {
	ctx := context.Background()

	//check
	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	//convert to string
	b, _ := json.Marshal(*newProduct)
	newProductString := string(b)
	newIdProductString := fmt.Sprintf("%d", newProduct.ID)

	//add
	_, err := pm.esClient.Index().
		Index(indexName).
		BodyJson(newProductString).
		Id(newIdProductString).
		Do(ctx)

	// call to flush data to disk for search. if no call --> need to wait for 5s to search since inserted
	pm.esClient.Refresh(indexName).Do(ctx)

	return err
}

func (pm *ProductManager) DeleteBook(product *database.Product) error {
	ctx := context.Background()

	//check
	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	//convert to string
	idProductString := fmt.Sprintf("%d", product.ID)

	//delete
	res, err := pm.esClient.Delete().
		Index(indexName).
		Id(idProductString).
		Do(ctx)

	if res.Shards.Successful > 0 {
		fmt.Println("Document deleted from from index: products, its id is :", idProductString)
	}

	return err
}
