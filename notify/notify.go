package notify

import (
	yahooapi "dashboard/yahooApi"
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
)

type GetSymbolList interface {
	Symbols() ([]yahooapi.Symbol, error)
	NotificationSymbols() ([]yahooapi.Symbol, error)
}

func StockNotification(chart yahooapi.YahooChart) (err error) {
	stockInfo := chart.Result[0].Metadata
	dayResult := stockInfo.RegMarketPrice - stockInfo.PreviousClose
	dayPercentage := dayResult / stockInfo.PreviousClose

	var dayResultString string
	var percentageString string
	if dayResult < 0 {
		dayResultString = fmt.Sprintf("<b>%.2f</b>", dayResult)
		percentageString = fmt.Sprintf("%.2f%%", dayPercentage*100)
	} else {
		dayResultString = fmt.Sprintf("<b>+%.2f</b>", dayResult)
		percentageString = fmt.Sprintf("+%.2f%%", dayPercentage*100)
	}

	beeep.Notify(
		"Stock Day End",
		fmt.Sprintf(
			"\n%s\t\t%.2f\t\t%s\t\t%s",
			chart.Result[0].Metadata.Symbol,
			stockInfo.RegMarketPrice,
			dayResultString,
			percentageString,
		),
		"",
	)

	return
}

func RunDaily(list GetSymbolList) (err error) {
	lastCheck := time.Now().Unix()
	for true {
		notiSymbols, err := list.NotificationSymbols()
		if err != nil {
			log.Println(err)
		}

		for _, stock := range notiSymbols {
			chart, err := yahooapi.GetStockData(stock)
			if err != nil {
				return err
			}

			stockInfo := chart.Result[0].Metadata
			dayEnd := stockInfo.TradingPeriods[0][0].End
			marketTime := stockInfo.RegMarketTime
			if marketTime >= dayEnd-100 && dayEnd != lastCheck {
				lastCheck = dayEnd
			}

			if err != nil {
				log.Println("Could not parse JSON")
				return err
			}
		}

		time.Sleep(time.Minute * 10)
	}

	return
}
