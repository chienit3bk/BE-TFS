package main

import (
	"crawdata/crawler"
	"crawdata/handleData"
	"crawdata/handleFile"
	"encoding/json"

	"fmt"
)

func main() {
	arr := crawler.Craw("https://www.sony.com.vn/electronics/tv/t/tv")
	arr2 := handleData.HandleData(arr)
	data, _ := json.Marshal(arr2)
	fmt.Println(string(data))
	handleFile.WriteToFile("fulldata.txt", string(data))
}
