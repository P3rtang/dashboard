package cmd

import (
	"dashboard/display"
	yahooapi "dashboard/yahooApi"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func ListTracked(cmd *cobra.Command, _ []string) {
	resp, err := http.Get("http://127.0.0.1:7444/api/tracked")
	if err != nil || resp.StatusCode >= 300 {
		log.Fatal(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var list []yahooapi.Symbol
	err = json.Unmarshal(body, &list)
	if err != nil {
		log.Fatal(err)
		return
	}

	display.DisplayTracked(list)
}
