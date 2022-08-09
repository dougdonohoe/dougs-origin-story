package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadSgbWords(t *testing.T) {
	length := len(SgbWords)
	assert.Equal(t, 5757, length)
	assert.Equal(t, "which", SgbWords[0])
	assert.Equal(t, "pupal", SgbWords[length-1])
	for _, word := range SgbWords {
		assert.Equal(t, 5, len(word))
	}
}

func TestLoadWordleWords(t *testing.T) {
	length := len(WordleWords)
	assert.Equal(t, 12974, length)
	assert.Equal(t, "aahed", WordleWords[0])
	assert.Equal(t, "zymic", WordleWords[length-1])
	for _, word := range WordleWords {
		assert.Equal(t, 5, len(word))
	}
}

func TestLoadScrabbleWords(t *testing.T) {
	length := len(ScrabbleWords)
	assert.Equal(t, 8636, length)
	assert.Equal(t, "aahed", ScrabbleWords[0])
	assert.Equal(t, "zymes", ScrabbleWords[length-1])
	for _, word := range ScrabbleWords {
		assert.Equal(t, 5, len(word))
	}
}

func TestFindBest(t *testing.T) {
	words := []string{"biiii", "zzizz", "yyiyo", "qqiqq"}

	word, value := FindBest(words, 'b', 0)
	assert.Equal(t, "biiii", word)
	assert.Equal(t, 7, value)

	word, value = FindBest(words, 'i', 2)
	assert.Equal(t, "zzizz", word)
	assert.Equal(t, 41, value)

	word, value = FindBest(words, 'o', 4)
	assert.Equal(t, "yyiyo", word)
	assert.Equal(t, 14, value)
}
