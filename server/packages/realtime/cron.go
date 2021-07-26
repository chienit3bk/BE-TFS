package realtime

import (
	"fmt"
	"project/packages/handlers/order"
	"project/packages/mail"
	"time"

	"github.com/robfig/cron/v3"
)

func SendEmailDailyReportForStoreOwner() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("30 16 18 * * *", createAndSendReport)
	c.Start()

	//////chạy realtime thì bỏ đi
	time.Sleep(time.Second * 600)
	c.Stop()
	// /////
}

func createAndSendReport() {
	//Get analytical data
	data := order.Statistic()

	//edit content
	statisticRevenue := fmt.Sprintf("Doanh thu: %v vnđ.\n", data.Revenue)
	statisticOrder := fmt.Sprintf("Số lượng đơn hàng mới trong 24 giờ qua: %v đơn hàng. \n", data.TotalOrder)
	statisticProductCount := fmt.Sprintf("Số lượng sản phẩm bán ra: %v sản phẩm. \n", data.TotalProduct)
	statisticProduct := "Top sản phẩm bán chạy: \n"
	for i, value := range data.TopSell {
		if i >= 5 {
			break
		}
		statisticProduct += fmt.Sprintf("Top %v: (%v sản phẩm): %v. \n", i+1, value.ProductCount, value.Product.Name)
	}
	reportContent := statisticRevenue + statisticOrder + statisticProductCount + statisticProduct

	//send email
	mail.SendReport(reportContent)
}

/*
	*  .  *  .  *  .  *  .  *  .  *
	-     -     -     -     -     -
	|     |     |     |     |     |
	|     |     |     |     |     +----- năm
	|     |     |     |     +----- tháng(1-12)
	|     |     |     +------- ngày(0-31)
	|     |     +--------- giờ(0-23)
	|     +----------- phút(0-59)
	+------------- giây(0-59)


// nếu một dấu * cần truyền nhiều giá trị thì dùng dấu / : vd : 0/30/45
*/
