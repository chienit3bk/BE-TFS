package elastic_search

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

type ESClient struct {
	*elastic.Client
}

func NewESClient(url string) (*ESClient, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
		
	fmt.Println("ES initialized...")
	return &ESClient{client}, err
}
