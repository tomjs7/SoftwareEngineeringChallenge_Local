package main

import (
	"fmt"
)

func main() {
	// itemsOrdered is a dummy value
	itemsOrdered := 12001
	packs, err := getPackSizes()

	if err != nil {
		fmt.Println("Error encountered while trying to get pack sizes:", err)
	}

	sortedPackSizes := sortPackSizes(packs.PackSizes)
	possiblePackSizes := possiblePackCombinations(sortedPackSizes)
	combinations := findCombinations(sortedPackSizes, possiblePackSizes)

	packsRequired := handlePacks(itemsOrdered, sortedPackSizes, possiblePackSizes, combinations)

	fmt.Print("\n")
	fmt.Println("Packs Required:", packsRequired)
	fmt.Print("\n")
}
