package main

import (
	"fmt"
	"github.com/ALMaclaine/ethplorer/ethplorer"
)

func main() {
	last, err := ethplorer.GetAddressTransactions(
		"0x68b22215ff74e3606bd5e6c1de8c2d68180c85f7",
		ethplorer.GetAddressTransactionsParams{Limit: 2},
		"")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(last)
}
