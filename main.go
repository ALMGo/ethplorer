package main

import (
	"fmt"
	"github.com/ALMaclaine/ethplorer/ethplorer"
)

func main() {
	last, err := ethplorer.GetTokensNew()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(last)
}
