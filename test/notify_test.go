package test_notify

import (
	"dashboard/notify"
	yahooapi "dashboard/yahooApi"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

type mockServer struct {
	symbols []yahooapi.Symbol
}

func newMockServer() (server *mockServer) {
	server = &mockServer{}
	server.symbols = []yahooapi.Symbol{
		{
			Exchange:       "TEST",
			ShortName:      "Mock",
			LongName:       "Mock Symbol",
			InstrumentType: "TEST",
			Symbol:         "AAPL",
			Score:          69,
			DoNotify:       true,
		},
	}
	return
}

func (self *mockServer) Symbols() (symbols []yahooapi.Symbol, err error) {
	return self.symbols, nil
}

func (self *mockServer) NotificationSymbols() (symbols []yahooapi.Symbol, err error) {
	return self.symbols, nil
}

func TestStockNotification(t *testing.T) {
	file, err := os.Open("./data/yahoo_chart.json")
	if err != nil {
		fmt.Println("Make sure test/data/yahoo_chart.json is available")
		t.Fatal(err)
	}
	bts, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	var chart yahooapi.YahooChart
	err = json.Unmarshal(bts, &chart)
	if err != nil {
		t.Fatal(err)
	}

	err = notify.LinuxStockNoti(&chart)
	if err != nil {
		t.Error(err)
	}

	return
}
