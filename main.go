package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var (
	wordle   = flag.Bool("wordle", false, "Use Wordle Words")
	scrabble = flag.Bool("scrabble", false, "Use Scrabble Words")
)

const (
	RED     = "\033[31m"
	RESTORE = "\033[0m"
)

// Find the highest value words which would complete the rows in this grid, where the
// value of the word is calculated using Scrabble letter values.  The columns need
// not form valid words.
//
//   |---|---|---|---|---|
//   | B |   |   |   |   |
//   |---|---|---|---|---|
//   |   | I |   |   |   |
//   |---|---|---|---|---|
//   |   |   | N |   |   |
//   |---|---|---|---|---|
//   |   |   |   | G |   |
//   |---|---|---|---|---|
//   |   |   |   |   | O |
//   |---|---|---|---|---|
//
// Word list defaults to Donald Knuth's Stanford GraphBase words (5757 words).
// The '-wordle' option switches to the list of all possible Wordle words (12974 words).
// The '-scrabble' option switches to the list of all possible Scrabble words (8636 words).
//
func main() {
	flag.Parse()

	if *wordle && *scrabble {
		fmt.Printf("Please use only use one of -scrabble or -wordle.\n")
		os.Exit(1)
	}

	words := SgbWords
	name := "Standford GraphBase"
	if *wordle {
		words = WordleWords
		name = "Wordle"
	} else if *scrabble {
		words = ScrabbleWords
		name = "Scrabble"
	}
	fmt.Printf("Using %d %s words (of which %d match BINGO)...\n",
		len(words), name, bingo(words))

	total := 0
	total += PrintBest(words, 'b', 0)
	total += PrintBest(words, 'i', 1)
	total += PrintBest(words, 'n', 2)
	total += PrintBest(words, 'g', 3)
	total += PrintBest(words, 'o', 4)
	fmt.Printf("  -----------\n")
	fmt.Printf("  TOTAL - %3d\n", total)
}

// PrintBest finds the best word with the given letter at the given place
// and then prints it out
func PrintBest(words []string, letter rune, place int) int {
	word, score := FindBest(words, letter, place)
	upper := fmt.Sprintf("%s%s%c%s%s",
		strings.ToUpper(word[0:place]),
		RED,
		unicode.ToUpper(rune(word[place])),
		RESTORE,
		strings.ToUpper(word[place+1:]))
	fmt.Printf("  %s - %3d\n", upper, score)
	return score
}

// bingo returns number of words that match BINGO criteria in the given list
func bingo(words []string) int {
	count := 0
	for _, word := range words {
		if rune(word[0]) == 'b' ||
			rune(word[1]) == 'i' ||
			rune(word[2]) == 'n' ||
			rune(word[3]) == 'g' ||
			rune(word[4]) == 'o' {
			count++
		}
	}
	return count
}
