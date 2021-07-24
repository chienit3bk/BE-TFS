package main 
import(
	"project/server"
	"project/packages/scheduler"
)
func main() {
	go scheduler.SendEmailDailyReportForStoreOwner()
	server.RunServer()
	
}