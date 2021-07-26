package realtime

import (
	"context"
	"encoding/json"
	"fmt"
	"project/database"
	"project/packages/handleFile"
	"project/packages/rabbitMQ/producer"
	"project/packages/rabbitMQ/rabbitmq"
	"strconv"
	"strings"
	"sync"
)

//create channel contain the new file name
var NewFileName = make(chan string, 5)

//realtime check : Has a new file uploaded ?
func CheckNewFile() {
	var newFile string
	for {
		newFile = <-NewFileName
		fmt.Println(newFile)
		handleNewFileUpload(newFile)
		saveProduct()
	}
}

// Đọc file từng dòng và chuyển chúng thành product,
// sau đó chuyển về dạng json và gửi lên rabbitMQ
// Đọc xong xóa file
func handleNewFileUpload(fileName string) {
	var line = make(chan string)
	var product = make(chan string)

	// đọc file, đọc xong sẽ xóa file
	go handleFile.ReadFileLineByLine(fileName, line)

	//handle line by line
	go func() {
		var fileLine string
		for {
			//lấy dữ liệu của một dòng
			fileLine = <-line
			if fileLine == "READDONE" { //đã đọc xong
				// gửi cho producer biết đã hết product để gửi lên
				product <- "DONE"
				break
			}

			// lấy ra newProduct
			newProduct, err := readLine(fileLine)
			if err != nil {
				continue
			}
			// chuyển dữ liệu cho producer
			data, _ := json.Marshal(newProduct)
			product <- string(data)
		}
	}()

	//check stop producer
	var tmp = make(chan string)
	var newProductString string
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for {
			newProductString = <-product
			//nếu dữ liệu được chuyển cho producer mang giá trị "DONE" thì dừng producer
			if newProductString == "DONE" {
				cancelFunc()
			} else { //nếu không thì chuyển cho producer để gửi
				tmp <- newProductString
				//đếm số product được chuyển đi
				productCounterOfSender++
			}
		}
	}()

	//send to rabbitMQ
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// khởi tọa rabbitmq
		rmq := rabbitmq.CreateNewRMQ(URI)

		pCh, err := rmq.GetChannel()
		if err != nil {
			fmt.Println("Cannot not get channel for producer")
			return
		}
		// khởi tạo producer
		producer := producer.CreateNewProducer(exchangeName, exchangeType, routingKey, pCh, ctx, &wg)

		// gửi dữ liệu đi
		producer.Send(tmp)
	}()

	wg.Wait()
}

// chuyển các thông tin trong một dòng thành một product hoàn chỉnh
func readLine(line string) (database.Product, error) {
	var product database.Product

	// chia nhỏ ra các trường thông tin chính
	fullInfo := strings.Split(line, ",")
	l := len(fullInfo)

	// kiểm tra đủ số trường thông tin tối thiểu
	if l > 10 {
		product.Name = fullInfo[0]
		product.LinkDetail = fullInfo[1]
		product.Technology = fullInfo[2]
		product.Resolution = fullInfo[3]
		product.Type = fullInfo[4]

		//get List Options
		listQuantitysString := strings.Split(fullInfo[l-1], "&&&")
		listSalePricesString := strings.Split(fullInfo[l-2], "&&&")
		listPricesString := strings.Split(fullInfo[l-3], "&&&")
		listSizesString := strings.Split(fullInfo[l-4], "&&&")
		for key, value := range listSizesString {
			priceInt, _ := strconv.Atoi(listPricesString[key])
			salePriceInt, _ := strconv.Atoi(listSalePricesString[key])
			quantityInt, _ := strconv.Atoi(listQuantitysString[key])

			product.ListOptions = append(product.ListOptions, database.Option{
				Size:      value,
				Price:     priceInt,
				SalePrice: salePriceInt,
				Quantity:  quantityInt,
			})
		}

		//get List Images
		listImagesString := strings.Split(fullInfo[l-5], "&&&")
		for _, link := range listImagesString {
			product.ListImages = append(product.ListImages, database.Image{
				Link: link,
			})
		}

		//get List Descriptions
		var fullDescription string
		for i := 5; i < l-5; i++ {
			fullDescription += fullInfo[i]
		}
		listDescriptionsString := strings.Split(fullDescription, "&&&")
		for _, description := range listDescriptionsString {
			product.ListDescriptions = append(product.ListDescriptions, database.Description{
				Content: description,
			})
		}
		return product, nil
	}
	return product, fmt.Errorf("cannot parse this line")
}
