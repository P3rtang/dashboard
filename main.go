package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StockJson struct {
	metadata  string `json:"Meta Data"`
	DailyData string `json:"Time Series (Daily)"`
}

func main() {
	resp, _ := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=IBM&outputsize=compact&datatype=json&apikey=442UJVP0RBHY7W1P")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Could not read response body")
		return
	}

	var stockData StockJson
	err = json.Unmarshal(body, &stockData)

	if err != nil {
		log.Println("Could not parse JSON")
		return
	}

	fmt.Println(stockData.metadata)
	// beeep.Notify("Title", "Message", "")
}
