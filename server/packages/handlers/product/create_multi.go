package product

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/database"
	es "project/packages/elastic_search"
	"project/packages/handlers/response"
	"project/packages/mail"
	"project/packages/rabbitMQ"
	"strconv"
	"strings"
	"sync"
)

// chưa thực hiện lấy được file ở client gửi về, chỉ mới  chạy được ở file local
func CreateMulti(w http.ResponseWriter, r *http.Request) {
	//read body with Claim: Content-Type: application/json
	//code

	//parse file
	//code

	// init rabbitMQ
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	producer, consumer, err := rabbitMQ.QuickCreateNewPairProducerAndConsumer("exc1", "queue1", ctx, &wg)
	if err != nil {
		response.ResponseWithJson(w, 500, map[string]string{"message": "Cannot create new pair croducer and consumer"})
		cancelFunc()
		return
	}

	//response
	response.ResponseWithJson(w, http.StatusCreated,
		map[string]string{"message": "The file is being processed, we will send an email notification when the processing is complete"})

	//create channel
	var sender = make(chan string)
	var receiver = make(chan string)

	//init msg counter for receiver
	counter := 0

	//read file Example (test local)
	f, err := os.Open("products.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(f)
	lines, err2 := reader.ReadAll()
	if err2 != nil {
		log.Fatal("Unable to parse file as CSV for products.csv", err)
	}
	go func() {
		for _, line := range lines {
			//get listDescriptions
			listDescriptionsString := strings.Split(line[5], "&&&")
			var listDescriptions []database.Description
			for _, descriptionString := range listDescriptionsString {
				listDescriptions = append(listDescriptions, database.Description{
					Content: descriptionString,
				})
			}

			//get listImages
			listImagesString := strings.Split(line[6], "&&&")
			var listImages []database.Image
			for _, imageString := range listImagesString {
				listImages = append(listImages, database.Image{
					Link: imageString,
				})
			}
			//get listOptions
			var listOptions []database.Option
			listSizesString := strings.Split(line[7], "&&&")
			listPricesString := strings.Split(line[8], "&&&")
			listSalePricesString := strings.Split(line[9], "&&&")
			listQuantityString := strings.Split(line[10], "&&&")

			for i, sizeString := range listSizesString {
				priceInt, _ := strconv.Atoi(listPricesString[i])
				salePriceInt, _ := strconv.Atoi(listSalePricesString[i])
				quantityInt, _ := strconv.Atoi(listQuantityString[i])
				listOptions = append(listOptions, database.Option{
					Size:      sizeString,
					Price:     priceInt,
					SalePrice: salePriceInt,
					Quantity:  quantityInt,
				})
			}
			var newProduct = database.Product{
				Name:             line[0],
				LinkDetail:       line[1],
				Technology:       line[2],
				Resolution:       line[3],
				Type:             line[4],
				ListDescriptions: listDescriptions,
				ListImages:       listImages,
				ListOptions:      listOptions,
			}
			productByte, _ := json.Marshal(newProduct)

			sender <- string(productByte)
		}
	}()

	wg.Add(3)
	//send to rabbitMQ
	go producer.Send(sender)
	go consumer.StartReceiveData(receiver)

	//handle data received
	var productString string
	db := database.ConnectToDatabase()
	go func() {
		for {
			productString = <-receiver
			var newProduct database.Product
			_ = json.Unmarshal([]byte(productString), &newProduct)
			//save product to database
			db.Create(&newProduct)

			//create variants
			var product2 database.Product
			db.Where("name = ?", newProduct.Name).Last(&product2)
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
			counter++
		}
	}()

	//check stop
	go func() {
		for {
			//số lượng nhận về bằng số lượng gửi đi
			if counter == len(lines) {
				cancelFunc()
				wg.Done()
				//sync data to Elastic search
				es.AutoSync()
				//send mail
				mail.SendNoticeImportSuccessful("ngoc nguyen", "nguyendinhhdpv3@gmail.com")
				break
			}
		}
	}()

	wg.Wait()

}
