package cmd

import (
	"bytes"
	yahooapi "dashboard/yahooApi"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func TrackStock(cmd *cobra.Command, args []string) {
	symbol := args[0]

	_, err := yahooapi.GetStockData(yahooapi.Symbol{Symbol: symbol})
	if err != nil {
		log.Println("Symbol Not Found")
		return
	}

	symbols, err := yahooapi.SearchStockSymbol(symbol)
	if err != nil {
		log.Println("Failed to search for symbol Info")
		return
	}

	symbolJson, err := json.Marshal(symbols.Quotes[0])
	if err != nil {
		log.Println("Failed to serialize symbol to JSON")
		return
	}

	reader := bytes.NewReader(symbolJson)
	_, err = http.Post("http://127.0.0.1:7444/api/add", "application/json", reader)
	if err != nil {
		log.Println(err)
		return
	}
}
