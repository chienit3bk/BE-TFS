package main

import (
	"project/packages/realtime"
	"project/server"
)

func main() {
	go realtime.SendEmailDailyReportForStoreOwner()
	go realtime.CheckNewFile()
	server.RunServer()

}
