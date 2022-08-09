package main

// Scrabble letter values from https://scrabble.hasbro.com/en-us/faq
//
//   (1 point)-A, E, I, O, U, L, N, S, T, R
//   (2 points)-D, G
//   (3 points)-B, C, M, P
//   (4 points)-F, H, V, W, Y
//   (5 points)-K
//   (8 points)-J, X
//   (10 points)-Q, Z
//
// Store in array indexed by ASCII value (wastes a little space but
// arguably a little faster than a map).
var letterValues = [123]byte{
	'a': 1,
	'e': 1,
	'i': 1,
	'o': 1,
	'u': 1,
	'l': 1,
	'n': 1,
	's': 1,
	't': 1,
	'r': 1,
	'd': 2,
	'g': 2,
	'b': 3,
	'c': 3,
	'm': 3,
	'p': 3,
	'f': 4,
	'h': 4,
	'v': 4,
	'w': 4,
	'y': 4,
	'k': 5,
	'j': 8,
	'x': 8,
	'q': 10,
	'z': 10,
}

// Score sums up letter values in given word and returns total
func Score(word string) int {
	total := 0
	for _, c := range word {
		total += int(letterValues[c])
	}
	return total
}
