package yahooapi

import "fmt"

type InstrumentType string

const (
	EQUITY         InstrumentType = "EQUITY"
	CRYPTOCURRENCY                = "CRYPTOCURRENCY"
	ETF                           = "ETF"
	INDEX                         = "INDEX"
)

type YahooJSON struct {
	Chart YahooChart `json:"chart"`
}

type YahooChart struct {
	Result []YahooResult `json:"result"`
	Error  *YahooError   `json:"error"`
}

type YahooError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
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

type YahooSymbolSearch struct {
	Quotes []Symbol `json:"quotes"`
}

type Symbol struct {
	Exchange       string         `json:"exchange"`
	ShortName      string         `json:"shortname"`
	LongName       string         `json:"longname"`
	InstrumentType InstrumentType `json:"quoteType"`
	Symbol         string         `json:"symbol"`
	Score          float64        `json:"score"`

	DoNotify bool
}

func (self *Symbol) Repr() (repr string) {
	repr += fmt.Sprintf("%-*s", 18, self.Symbol)
	repr += fmt.Sprintf("%-*s", 12, self.Exchange)
	repr += fmt.Sprintf("%-*s", 18, self.ShortName)

	return
}
