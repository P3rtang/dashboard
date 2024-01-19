package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func ToggleNotify(cmd *cobra.Command, args []string) {
	symbol := args[0]

	if cmd.Flag("remove").Value.String() == "true" {
		err := DisableNotify(symbol)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := EnableNotify(symbol)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func EnableNotify(symbol string) (err error) {
	resp, err := http.Post("http://127.0.0.1:7444/api/notify/"+symbol, "application/json", nil)

	if err != nil {
		return
	}

	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
	}

	return
}

func DisableNotify(symbol string) (err error) {
	resp, err := http.Post("http://127.0.0.1:7444/api/denotify/"+symbol, "application/json", nil)

	if err != nil {
		return
	}

	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
	}

	return
}
