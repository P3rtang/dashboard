package notify

import (
	yahooapi "dashboard/yahooApi"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
)

type GetSymbolList interface {
	Symbols() ([]yahooapi.Symbol, error)
	NotificationSymbols() ([]yahooapi.Symbol, error)
}

func StockNotification(chart *yahooapi.YahooChart) (err error) {
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
			"\n%s        %.2f      %s  (%s)",
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
	var lastCheck int64 = 0
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
			if chart == nil || len(chart.Result) == 0 {
				return errors.New("Unable to fetch chart data")
			}

			stockInfo := chart.Result[0].Metadata
			dayEnd := stockInfo.TradingPeriods[0][0].End
			marketTime := stockInfo.RegMarketTime
			if marketTime >= dayEnd-100 && dayEnd >= lastCheck {
				StockNotification(chart)
			}

			if err != nil {
				log.Println("Could not parse JSON")
				return err
			}
		}

		lastCheck = time.Now().Unix()
		time.Sleep(time.Minute * 10)
	}

	return
}
