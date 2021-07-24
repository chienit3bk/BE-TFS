package main

import (
	"encoding/json"
	"ngoc/database"
	"strconv"
	"strings"
	"time"

	"ngoc/fileHandling"
	// csv "ngoc/create_csv_file"
)

//"79,900,000 VNĐ" to 79900000
func convertPrice(s string) int {
	s1 := strings.Split(s, " ")
	s2 := strings.Split(s1[0], ",")
	r := strings.Join(s2, "")

	Int, _ := strconv.Atoi(r)
	return Int
}

func main() {

	// csv.Create()  //chạy để tạo ra file csv

	s := fileHandling.ReadFile()

	var arr = make([]database.Tivi, 1)
	json.Unmarshal([]byte(s), &arr)
	db := database.ConnectToDatabase()
	// var idimg int = 1
	// var iddes int = 1
	var idop int = 1

	//import 41 tivi
	for _, tivi := range arr {
		var product = database.Product{
			// ID:          int(tivi.ID),
			Name:       tivi.Name,
			LinkDetail: tivi.LinkDetail,
			Technology: tivi.Technology,
			Resolution: tivi.Resolution,
			Type:       tivi.Type,
		}
		db.Create(&product)
		for _, description := range tivi.Description {
			var des = database.Description{
				// ID:         iddes,
				IDProduct: int(tivi.ID),
				Content:   description,
			}
			db.Create(&des)
			// iddes++
		}
		for i, size := range tivi.Sizes {
			var op = database.Option{
				ID:        idop,
				IDProduct: int(tivi.ID),
				Size:      size,
				Price:     convertPrice(tivi.Prices[i]),
				SalePrice: convertPrice(tivi.Prices[i]),
				Quantity:  20,
			}
			db.Create(&op)
			var variant = database.Variant{
				IDProduct:   int(tivi.ID),
				IDOption:    idop,
				ProductName: tivi.Name,
				Size:        size,
				Price:       convertPrice(tivi.Prices[i]),
				SalePrice:   convertPrice(tivi.Prices[i]),
			}
			db.Create(&variant)

			idop++
		}
		for _, link := range tivi.Imgs {
			var img = database.Image{
				// ID:         idimg,
				IDProduct: int(tivi.ID),
				Link:      link,
			}
			db.Create(&img)
			// idimg++
		}
		// 	// fmt.Println(tivi.ID)
		// 	// fmt.Println(tivi.Name)
		// 	// fmt.Println(tivi.Technology)
		// 	// fmt.Println(tivi.Resolution)
		// 	// fmt.Println(tivi.Type)
		// 	// fmt.Println(tivi.Imgs)
		// 	// fmt.Println(tivi.Description)
		// 	// fmt.Println(tivi.Sizes)
		// 	// fmt.Println(tivi.Prices)
		// 	// fmt.Println(tivi.LinkDetail)
		// 	// fmt.Println(arr[40])
	}

	var user = database.User{
		// ID:       1,
		Name:     "ngoc",
		Username: "ngochd",
		Password: "$2a$14$6HrcAfhpFwB1xMC2nqULf.1ZLf7YbFQLFLhlpoZJamh/AXYx9qTfS", //tương đương "ngochd246",
		Email:    "nguyenngochdpv3@gmail.com",
		Address:  "hanoi vietnam",
		Phone:    "0925933543",
		IsAdmin:  true,
	}
	db.Create(&user)

	var order = database.Order{
		// ID:           1,
		IDUser:      1,
		TotalPrice:  convertPrice(arr[0].Prices[0]) + convertPrice(arr[1].Prices[0]),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeleteAt:    time.Now(),
		Username:    "ngochd",
		PhoneUser:   "0925933543",
		AddressUser: "hanoi vietnam",
		State:       "paid",
	}
	db.Create(&order)

	var orderDetail1 = database.OrderDetail{
		// ID:           1,
		IDVariant:   1,
		IDOrder:     1,
		Quantity:    1,
		ProductName: arr[0].Name,
		TotalPrice:  convertPrice(arr[0].Prices[0]),
	}
	db.Create(&orderDetail1)
	var orderDetail2 = database.OrderDetail{
		// ID:           2,
		IDVariant:   2,
		IDOrder:     1,
		Quantity:    1,
		ProductName: arr[1].Name,
		TotalPrice:  convertPrice(arr[1].Prices[0]),
	}
	db.Create(&orderDetail2)

	var bill = database.Bill{
		// ID:           1,
		IDOrder:     1,
		IDUser:      1,
		IdAdmin:     1,
		Username:    "ngoc",
		AdminName:   "ngoc",
		PhoneUser:   "0925933543",
		TotalPrice:  convertPrice(arr[0].Prices[0]) + convertPrice(arr[1].Prices[0]),
		AddressUser: "hanoi vietnam",
		State:       "paid",
		CreatedAt:   time.Now(),
	}
	db.Create(&bill)

}
