package display

import (
	"dashboard/functions"
	yahooapi "dashboard/yahooApi"
	"fmt"
)

func DisplaySearch(symbols []yahooapi.Symbol) {
	var symbolPadding int
	var exchangePadding int
	var shortNamePadding int
	for _, symbol := range symbols {
		symbolPadding = functions.Max(symbolPadding, len(symbol.Symbol))
		exchangePadding = functions.Max(exchangePadding, len(symbol.Exchange))
		shortNamePadding = functions.Max(shortNamePadding, len(symbol.ShortName))
	}

	// the padding here is set 8 character too high to offset for the esacpe sequence
	fmt.Printf("%-*s", functions.Max(9, symbolPadding+3)+8, "\033[4mSymbol\033[0m")
	fmt.Printf("%-*s", functions.Max(11, exchangePadding+3)+8, "\033[4mExchange\033[0m")
	fmt.Printf("%-*s", functions.Max(12, shortNamePadding+3)+8, "\033[4mShortName\033[0m")
	fmt.Println()
	fmt.Println()

	for _, symbol := range symbols {
		fmt.Printf("%-*s", functions.Max(9, symbolPadding+3), symbol.Symbol)
		fmt.Printf("%-*s", functions.Max(11, exchangePadding+3), symbol.Exchange)
		fmt.Printf("%-*s", functions.Max(12, shortNamePadding+3), symbol.ShortName)
		fmt.Println()
	}
}

func DisplayTracked(symbols []yahooapi.Symbol) {
	var symbolPadding int
	var exchangePadding int
	var shortNamePadding int
	for _, symbol := range symbols {
		symbolPadding = functions.Max(symbolPadding, len(symbol.Symbol))
		exchangePadding = functions.Max(exchangePadding, len(symbol.Exchange))
		shortNamePadding = functions.Max(shortNamePadding, len(symbol.ShortName))
	}

	// the padding here is set 8 character too high to offset for the esacpe sequence
	fmt.Printf("%-*s", functions.Max(9, symbolPadding+3)+8, "\033[4mSymbol\033[0m")
	fmt.Printf("%-*s", functions.Max(11, exchangePadding+3)+8, "\033[4mExchange\033[0m")
	fmt.Printf("%-*s", functions.Max(12, shortNamePadding+3)+8, "\033[4mShortName\033[0m")
	fmt.Printf("%-*s", 9, "\033[4mNotify\033[0m")
	fmt.Println()
	fmt.Println()

	for _, symbol := range symbols {
		fmt.Printf("%-*s", functions.Max(9, symbolPadding+3), symbol.Symbol)
		fmt.Printf("%-*s", functions.Max(11, exchangePadding+3), symbol.Exchange)
		fmt.Printf("%-*s", functions.Max(12, shortNamePadding+3), symbol.ShortName)
		if symbol.DoNotify {
			fmt.Printf("%-*s", 9, "Yes")
		} else {
			fmt.Printf("%-*s", 9, "No")
		}
		fmt.Println()
	}
}
