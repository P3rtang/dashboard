package cmd

import (
	"dashboard/notify"
	server "dashboard/server"
	yahooapi "dashboard/yahooApi"
	"log"
	"runtime"

	"github.com/spf13/cobra"
)

var SYSTEM_NOTI_CALLBACK func(*yahooapi.YahooChart) (err error)

func init() {
	switch runtime.GOOS {
	case "linux":
		SYSTEM_NOTI_CALLBACK = notify.LinuxStockNoti
	case "windows":
		SYSTEM_NOTI_CALLBACK = notify.WindowsStockNoti
	}
}

func Serve(cmd *cobra.Command, _ []string) {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	go server.Serve()
	err = notify.RunDaily(server.Database, SYSTEM_NOTI_CALLBACK)
	if err != nil {
		log.Fatal(err)
	}
}
