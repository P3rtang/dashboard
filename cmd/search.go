package cmd

import (
	"dashboard/display"
	yahooapi "dashboard/yahooApi"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func SearchStock(cmd *cobra.Command, args []string) {
	result, err := yahooapi.SearchStockSymbol(args[0])
	if err != nil {
		log.Println(err)
		return
	}

	resultAmount, err := strconv.ParseInt(cmd.Flag("number").Value.String(), 10, 0)
	if err != nil {
		log.Println("Could not interpret number flag as a integer")
	}

	display.DisplaySearch(result.Quotes[0:resultAmount])
}
