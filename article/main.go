package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed words.txt
var bingoWords string

func main() {
	BingoWords := strings.Split(bingoWords, "\n")
	for _, word := range BingoWords {
		fmt.Printf("WORD    %s\n", word)
	}
}
