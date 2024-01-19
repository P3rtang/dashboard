package cmd

import (
	"dashboard/notify"
	server "dashboard/server"
	"log"

	"github.com/spf13/cobra"
)

func Serve(cmd *cobra.Command, _ []string) {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	go server.Serve()
	notify.RunDaily(server.Database)
}
