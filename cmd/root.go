package root

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

type InstrumentType string

const (
	EQUITY InstrumentType = "EQUITY"
)

type YahooJSON struct {
	Chart YahooChart `json:"chart"`
}

type YahooChart struct {
	Result []YahooResult `json:"result"`
	Error  string        `json:"error"`
}

type YahooResult struct {
	Metadata YahooMetadata `json:"meta"`
}

type YahooMetadata struct {
	Currency       string `json:"currency"`
	Symbol         string `json:"symbol"`
	Exchange       string `json:"exchangeName"`
	InstrumentType InstrumentType
	RegMarketPrice float64 `json:"regularMarketPrice"`
	RegMarketTime  int64   `json:"regularMarketTime"`
	PreviousClose  float64 `json:"previousClose"`

	TradingPeriods [][]TradingPeriod
}

type TradingPeriod struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	GmtOffset int64  `json:"gmtoffset"`
}

func RunDaily(_ *cobra.Command, _ []string) {
	lastCheck := time.Now().Unix()
	for true {
		resp, err := http.Get("https://query1.finance.yahoo.com/v8/finance/chart/AIR.PA?region=US&lang=en-US&includePrePost=false&interval=2m&useYfid=true&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance")

		if err != nil {
			log.Println("Could not fetch Stock data")
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Println("Could not read response body")
			return
		}

		var stockData YahooJSON
		err = json.Unmarshal(body, &stockData)

		if err != nil {
			log.Println("Could not parse JSON")
			return
		}

		stockInfo := stockData.Chart.Result[0].Metadata
		dayEnd := stockInfo.TradingPeriods[0][0].End
		marketTime := stockInfo.RegMarketTime

		if marketTime >= dayEnd-100 && dayEnd != lastCheck {
			lastCheck = dayEnd
			dayResult := stockInfo.RegMarketPrice - stockInfo.PreviousClose

			var dayResultString string
			if dayResult < 0 {
				dayResultString = fmt.Sprintf("<b>%f</b>", dayResult)
			} else {
				dayResultString = fmt.Sprintf("%f", dayResult)
			}

			beeep.Notify(
				"Stock Day End",
				fmt.Sprintf("\n%s\t\t%f\t%s",
					stockData.Chart.Result[0].Metadata.Symbol,
					stockInfo.RegMarketPrice,
					dayResultString,
				),
				"",
			)
		}

		time.Sleep(time.Minute * 10)
	}
}

var rootCmd = &cobra.Command{
	Use:   "Start the dashboard server",
	Run:   RunDaily,
	Short: "Personal stock dashboard",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
