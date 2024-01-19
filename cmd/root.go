package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Personal stock dashboard",
}

func init() {
	var searchCmd = &cobra.Command{
		Use:   "search <name>",
		Run:   SearchStock,
		Short: "Seach for a stock by name",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	}

	searchCmd.Flags().IntP("number", "n", 5, "Number of results returned (max 5)")
	rootCmd.AddCommand(searchCmd)

	var serveCmd = &cobra.Command{
		Use:   "serve",
		Run:   Serve,
		Short: "Start the notification server",
	}

	rootCmd.AddCommand(serveCmd)

	var trackCmd = &cobra.Command{
		Use:   "track",
		Run:   TrackStock,
		Short: "Keep track of a stock",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	}

	rootCmd.AddCommand(trackCmd)

	var listTrackedCmd = &cobra.Command{
		Use:   "ls",
		Run:   ListTracked,
		Short: "Show your tracked stocks",
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(listTrackedCmd)

	var notifyCmd = &cobra.Command{
		Use:   "notify <symbol>",
		Run:   ToggleNotify,
		Short: "Notify of price changes",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	}
	notifyCmd.Flags().BoolP("remove", "r", false, "Remove notification")

	rootCmd.AddCommand(notifyCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
