package main

import (
	"fmt"
	"github.com/ALMaclaine/ethplorer/ethplorer"
)

func main() {
	last, err := ethplorer.GetTokenDailyPriceHistory(
		"0xdf9d4674a430bdcc096a3a403128357ab36844ba",
		ethplorer.GetTokenHistoryGroupedParams{},
	"")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(last)
}
