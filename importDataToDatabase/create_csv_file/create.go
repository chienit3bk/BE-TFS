package create_csv_file

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"ngoc/database"
	"os"
	"strings"

	"ngoc/fileHandling"
)

//"79,900,000 VNĐ" to 79900000
func convertPriceToString(m []string) []string {
	mm := make([]string, 1)
	for _, v := range m {
		s1 := strings.Split(v, " ")
		s2 := strings.Split(s1[0], ",")
		r := strings.Join(s2, "")
		mm = append(mm, r)
	}
	return mm
}

func removeElementByIndex(s []string, index int) []string {
    ret := make([]string, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func Create() {
	s := fileHandling.ReadFile()

	var arr []database.Tivi
	json.Unmarshal([]byte(s), &arr)
	


	var m = make([][]string, len(arr))
	//import 41 tivi
	for i , tivi := range arr {

		var listDess = strings.Join(tivi.Description, "&&&")
		var listImgs = strings.Join(tivi.Imgs, "&&&")
		var listSizes = strings.Join(tivi.Sizes, "&&&") 
		tivi.Prices = removeElementByIndex(convertPriceToString(tivi.Prices), 0)
		var listPrices = strings.Join(tivi.Prices, "&&&")
		var listSalePrices = listPrices
		var listQuantitys = "20"
		for i := range tivi.Sizes {
			if i == 0 {
				continue
			}
			listQuantitys += "&&&20"
		}
		var m1 []string
		m1 = append(m1, tivi.Name, tivi.LinkDetail, tivi.Technology, tivi.Resolution, tivi.Type, listDess, 
					listImgs, listSizes, listPrices, listSalePrices, listQuantitys)
		m[i] = m1
	}
	// Load a csv file.
    f, err := os.Create("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(m[0])

    // Create a new writer
    w := csv.NewWriter(f)
	w.WriteAll(m)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	fmt.Println("create csv file successful !")
}
