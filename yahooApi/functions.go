package yahooapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func SearchStockSymbol(query string) (result *YahooSymbolSearch, err error) {
	url := fmt.Sprintf(
		"%s%s%s",
		"https://query1.finance.yahoo.com/v1/finance/search?q=",
		query,
		"&format=json&env=store://datatables.org/alltableswithkeys",
	)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}

func GetStockData(symbol Symbol) (result *YahooChart, err error) {
	url := fmt.Sprintf(
		"%s%s%s",
		"https://query1.finance.yahoo.com/v8/finance/chart/",
		symbol.Symbol,
		"?region=US&lang=en-US&includePrePost=false&interval=2m&useYfid=true&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance",
	)

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	var stockData YahooJSON
	err = json.Unmarshal(body, &stockData)

	if err != nil {
		return
	}

	if stockData.Chart.Error != nil {
		err = errors.New(stockData.Chart.Error.Code)
		return
	}

	return &stockData.Chart, nil
}
