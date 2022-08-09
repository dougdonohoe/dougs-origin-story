package main

import (
	_ "embed"
	"strings"
)

var (
	SgbWords      []string
	WordleWords   []string
	ScrabbleWords []string

	// Stanford GraphBase words
	// Src: https://www-cs-faculty.stanford.edu/~knuth/sgb.html
	//go:embed sgb-words.txt
	sgbWords string

	// Wordle all valid guesses
	// Src: Wordle source code by way of https://github.com/tabatkins/wordle-list
	//go:embed wordle.txt
	wordleWords string

	// Scrabble 5-letter words
	// Src: Denison University http://discovercs.denison.edu/chapter8/scrabble.txt
	//go:embed scrabble.txt
	scrabbleWords string
)

// init translates embedded word files into slice of strings
func init() {
	WordleWords = strings.Split(wordleWords, "\n")
	SgbWords = strings.Split(sgbWords, "\n")
	ScrabbleWords = strings.Split(scrabbleWords, "\n")
}

// FindBest finds the highest value word with a given letter at given spot
func FindBest(words []string, letter rune, place int) (string, int) {
	bestWord := ""
	bestValue := 0
	for _, word := range words {
		if rune(word[place]) != letter {
			continue
		}
		value := Score(word)
		if value > bestValue {
			bestWord = word
			bestValue = value
		}
	}
	return bestWord, bestValue
}
