package main

import (
	"fmt"
	"sort"
)

func main() {

	// data input
	var text string
	fmt.Scan(&text)

	// convert to runes
	textRunes := []rune(text)

	// create and fill a map with runes and their frequencies
	mapRunes := make(map[rune]int)
	for _, r := range textRunes {
		mapRunes[r]++
	}

	// get a list of runes and sort them
	runes := make([]rune, len(mapRunes))
	i := 0
	for r := range mapRunes {
		runes[i] = r
		i++
	}
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	// data output
	for _, r := range runes {
		fmt.Print(string(r), mapRunes[r])
	}
}
