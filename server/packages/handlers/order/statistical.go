package order

import (
	"net/http"
	"project/database"
	"project/packages/handlers/response"
	"strings"
	"time"
)

func StatisticsHourly(w http.ResponseWriter, r *http.Request) {
	//get data
	fullStatisticResult := Statistic()
	//response
	response.ResponseWithJson(w, http.StatusOK, fullStatisticResult)
}

type StatisticData struct {
	Revenue          int                `json:"revenue"`
	TotalOrder       int                `json:"total_order"`
	TotalProduct     int                `json:"total_product"`
	ListTimes        []string           `json:"list_times"`
	ListOrders       []int              `json:"list_orders"`
	ListProducts     []int              `json:"list_products"`
	MapStatistiOrder map[int]Duration   `json:"map_order_statistic"`
	TopSell          []ProductStatistic `json:"top_sell"`
}

type Duration struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	OrderCount   int       `json:"order_count"`
	ProductCount int       `json:"product_count"`
}

func (d Duration) inTimeSpan(t time.Time) bool {
	return t.After(d.Start) && t.Before(d.End)
}

type ProductStatistic struct {
	Product      database.Product `json:"product"`
	ProductCount int              `json:"product_count"`
}

func Statistic() StatisticData {
	timenow := time.Now()
	timestart := timenow.Add(-time.Hour * 24)
	timestart = rounding(timestart)

	//create map statistic order map[STT]Duration
	mapOrderStatistic := createMapStatisticOrder(timenow)

	//create map statistic product map[Product_name]Counter
	mapProductNameStatistic := make(map[string]int)

	revenue := 0
	totalProduct := 0

	//get orders
	var listOrders []database.Order
	db := database.ConnectToDatabase()
	db.Where("created_at BETWEEN ? AND ?", timestart, timenow).Find(&listOrders)

	//handle list orders
	for _, order := range listOrders {
		//update revenue
		revenue += order.TotalPrice

		//get listOrderDetail
		var listOrderDetails []database.OrderDetail
		db.Where("id_order = ?", order.ID).Find(&listOrderDetails)
		order.ListOrderDetails = listOrderDetails

		//create key: save index of duration
		var keyy = 0

		//update map statistic order
		for key, value := range mapOrderStatistic {
			if value.inTimeSpan(order.CreatedAt) {
				mapOrderStatistic[key] = Duration{
					Start:        value.Start,
					End:          value.End,
					OrderCount:   value.OrderCount + 1,
					ProductCount: value.ProductCount,
				}
				keyy = key
				break
			}
		}

		//update map statistic product
		for _, orderDetail := range order.ListOrderDetails {
			_, ok := mapProductNameStatistic[orderDetail.ProductName]
			if ok {
				mapProductNameStatistic[orderDetail.ProductName] += orderDetail.Quantity
			} else {
				mapProductNameStatistic[orderDetail.ProductName] = 1
			}
			//update totalProduct
			totalProduct += orderDetail.Quantity

			//update product count
			oldDuration := mapOrderStatistic[keyy]
			mapOrderStatistic[keyy] = Duration{
				Start:        oldDuration.Start,
				End:          oldDuration.End,
				OrderCount:   oldDuration.OrderCount,
				ProductCount: oldDuration.ProductCount + orderDetail.Quantity,
			}
		}

	}

	//convert map to slice statistic
	SliceProductStatistic := convertToSlice(mapProductNameStatistic)

	//convert data
	listTimess, listOrderss, listProductss := convertMapOrderStatistic(mapOrderStatistic)
	return StatisticData{
		Revenue:          revenue,
		TotalOrder:       len(listOrders),
		TotalProduct:     totalProduct,
		ListTimes:        listTimess,
		ListOrders:       listOrderss,
		ListProducts:     listProductss,
		MapStatistiOrder: mapOrderStatistic,
		TopSell:          SliceProductStatistic,
	}

}

//l??m tr??n th???i gian v??? gi???
func rounding(t time.Time) time.Time {
	y := t.Year()
	m := t.Month()
	d := t.Day()
	h := t.Hour()
	return time.Date(y, time.Month(m), d, h+1, 0, 0, 0, time.Now().UTC().Local().Location())
}

func convertMapOrderStatistic(m map[int]Duration) ([]string, []int, []int) {
	var (
		listTimes    = make([]string, 25)
		listOrders   = make([]int, 25)
		listProducts = make([]int, 25)
	)
	for key, duration := range m {
		listTimes[key] = getTime(duration.Start.String())
		listOrders[key] = duration.OrderCount
		listProducts[key] = duration.ProductCount
	}
	return listTimes, listOrders, listProducts
}

//"2021-07-25 20:00:00 +0700 +07 2021-07-25 21:00:00 +0700 +07 0 0" to "20:00"
func getTime(s string) (r string) {
	arr := strings.Split(s, " ")
	arr2 := strings.Split(arr[1], ":")
	r = arr2[0] + ":" + arr2[1]
	return
}

// T???o m???t map th???ng k?? v???i key l?? s??? th??? t??? s???p x???p theo th???i gian, value l?? c??c kho???ng th???i gian (1 ti???ng) v?? s???
// l?????ng order ???????c t???o ra trong kho???ng th???i gian ????
func createMapStatisticOrder(timenow time.Time) map[int]Duration {
	m := make(map[int]Duration)

	timeStart := timenow.Add(-time.Hour * 24)
	timeStart = rounding(timeStart)

	for i := 1; i <= 24; i++ {
		m[i] = Duration{
			Start:        timeStart,
			End:          timeStart.Add(time.Hour * 1),
			OrderCount:   0,
			ProductCount: 0,
		}
		timeStart = timeStart.Add(time.Hour * 1)
	}

	return m
}

//convert MapStatisticProduct To SliceStatisticProduct by order (arranged in order) and add more product info
func convertToSlice(mapProductNameStatistic map[string]int) (result []ProductStatistic) {
	//connect database
	db := database.ConnectToDatabase()

	//get length of mapProductNameStatistic
	l := len(mapProductNameStatistic)

	//update slice
	for i := 0; i < l; i++ {
		//find max count
		maxCount := 0
		nameProductMaxCount := ""
		for name, count := range mapProductNameStatistic {
			if count >= maxCount {
				nameProductMaxCount = name
				maxCount = count
			}
		}
		//delete max from mapProductNameStatistic
		delete(mapProductNameStatistic, nameProductMaxCount)

		//get full info
		var product database.Product
		db.Where("name = ?", nameProductMaxCount).First(&product)

		//slice append max value
		result = append(result, ProductStatistic{
			Product:      product,
			ProductCount: maxCount,
		})
	}

	return
}
