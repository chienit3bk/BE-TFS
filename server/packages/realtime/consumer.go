package realtime

import (
	"context"
	"encoding/json"
	"fmt"
	"project/database"

	"project/packages/mail"
	"project/packages/rabbitMQ/consumer"
	"project/packages/rabbitMQ/rabbitmq"
	"sync"
)

func saveProduct() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)

	//khởi tạo rabbitmq
	rmq := rabbitmq.CreateNewRMQ(URI)

	cCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel for consumer")
		cancelFunc()
		return
	}

	//khởi tạo consumer
	consumer := consumer.CreateNewConsumer(exchangeName, exchangeType, bindingKey, queueName, cCh, ctx, &wg)

	// khởi tạo channel nhận dữ liệu
	productReceiver := make(chan string)

	//get product
	go consumer.StartReceiveData(productReceiver)

	//xử lý dữ liệu nhận về và lưu vào database
	var (
		newProduct    database.Product
		productString string
	)
	db := database.ConnectToDatabase()
	go func() {
		for {
			productString = <-productReceiver

			productCounterOfReceiver++
			_ = json.Unmarshal([]byte(productString), &newProduct)

			//create product
			db.Create(&newProduct)

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
		}
	}()

	//check Stop
	go func() {
		for {
			//nếu số sản phẩm nhận về bằng số sản phẩm gửi lên thì dừng consumer lại
			if productCounterOfSender != 0 && productCounterOfReceiver == productCounterOfSender {
				productCounterOfSender = 0
				productCounterOfReceiver = 0

				// gửi mail đã xử lý xong file
				mail.SendNoticeImportSuccessful("ngoc nguyen", "nguyendinhhdpv3@gmail.com")
				cancelFunc()
			}
		}
	}()
	wg.Wait()
}
