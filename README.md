# 80s Word Puzzles and the Atari 800 - My Origin Story

This repository is a companion to my [Medium post](https://medium.com/@DougDonohoe/80s-word-puzzles-and-the-atari-800-my-origin-story-3c86fa5947dc) 
discussing the first serious computer program I wrote, as a teenager in the 1980s, 
to solve a mail-in word puzzle using an Atari 800.  Please see that post for 
more background on this code.

## The Puzzle

The problem to solve is to find the highest value words that can fill in 
the rows of this grid, where the value of the word is calculated by summing each letter's
Scrabble letter value.  The columns need not form valid words.

```text
|---|---|---|---|---|
| B |   |   |   |   |
|---|---|---|---|---|
|   | I |   |   |   |
|---|---|---|---|---|
|   |   | N |   |   |
|---|---|---|---|---|
|   |   |   | G |   |
|---|---|---|---|---|
|   |   |   |   | O |
|---|---|---|---|---|
```

## Explanation

This is a quasi-reproduction of my original code, which was written in BASIC 
over 40 years ago for the Atari 800, but re-imagined in Go.  Why Go?  Because
it is what I'm currently working in professionally, and I somehow feel like it
carries the ethos of BASIC today in that it is also general purpose, high level
and easy to use.  Also, because I didn't want to re-learn 
[Atari BASIC](https://en.wikipedia.org/wiki/Atari_BASIC) and re-implement
the solution in it, although I [partially did](atari/README.md).

Back then, it would be over 15 years before unit tests were invented.  I would
have had no test code.  I have a small amount of it this repo as unit testing has since
gotten into my blood, even for simple code like this.  I imagine my original
program was one big file, including the list of words.  I've broken things up
a bit, including putting the word lists in their own files.  However, I have
not endeavored to reduce all inefficiencies or make this code perfect.
Can code ever be perfect?  Perfection is in the eye of the beholder.  Sure,
I could cache word scores or create a `struct` to encapsulate a word list and
its name, but I'm not getting graded on this (except maybe in Hacker News).

Back then, there would have been one list of words, manually sourced from
a Merriam-Webster Dictionary or similar.  For this reproduction, I have
found three lists of 5-letter words to chose from.

The first (and default) option, is the 5,757 words from Donald Knuth's Stanford 
GraphBase book which seems like a very generous estimate of how many 5-letter words 
would have been in a dictionary in the 1980s.  I can't imagine my family
would have typed in more than five thousand words!

The second list, using `-wordle`, is from (you guessed it), the list of 12,974 
valid guesses for the popular Wordle word game.

The third list, using `-scrabble`, is all five-letter Scrabble words,
or at least what is found in (coincidentally), something my alma-mater posted in an
[on-line class assignment list](http://discovercs.denison.edu/chapter8/).  I scraped
out all 8,636 5-letter words.

The word letter values used back then for the contest are lost to time, so I've
used the official Scrabble letter values to calculate the score on the assumption
that the contest would have assigned higher values to less frequent letters.

My solution sticks with the first word it finds with the highest score.  There
may be other words that score the same, but for the purposes of the contest,
that seems irrelevant. 

## The Code

See `score.go` for the code that calculates the score of a word.

See `words.go` for the code that loads words from a text file and the code
that finds the best word for a letter at a given position.

See `main.go` for the code that pulls it all together.

## Running it

You need Go installed, which on a Mac can be done with `brew`:

```shell
brew install go
```

For full installation instructions, including other platforms, please
see the [official Go docs](https://go.dev/doc/install).

Once you have Go installed, change to the root of this repository and run:

```shell
$ go run .
Using 5757 Standford GraphBase words (of which 1525 match BINGO)...
  BUZZY -  28
  FIZZY -  29
  JUNKY -  19
  ZINGY -  18
  MEZZO -  25
  -----------
  TOTAL - 119
```

To use the Wordle word list:

```shell
$ go run . -wordle
Using 12974 Wordle words (of which 3381 match BINGO)...
  BEZZY -  28
  FIZZY -  29
  ZANZA -  23
  AZYGY -  21
  BIZZO -  25
  -----------
  TOTAL - 126
```

To use the Scrabble word list:

```shell
$ go run . -scrabble
Using 8636 Scrabble words (of which 2219 match BINGO)...
  BOOZY -  19
  FIZZY -  29
  ZANZA -  23
  ZINGY -  18
  MEZZO -  25
  -----------
  TOTAL - 114
```

## Testing it

To run all unit tests:

```shell
go test . -v
```

## Random Thoughts

This problem isn't too difficult, really, at least today, using modern tools.
I think it was probably decently difficult to 13-year-old me using Atari BASIC.
Actually the hardest part, even now, was sourcing the right list of 5-letter words.

Even though I rather hate interview coding challenges, I actually think it would 
make for a half-way decent one for an 45 or 60 minute interview, provided one 
provides the code to access the list of words ahead of time (honestly, who knows 
`go:embed` or the proper way to read and parse a file in their favorite programming
language off the top of their head?).  Depending on the language, you test knowledge of 
arrays and looping and characters and strings and maybe even maps for optimization.
Solving this doesn't require remembering the exact right algorithm or 'trick', just
basic programming 101.

To any future interviewees that encounter this as a question, I wish you the best of luck,
and I apologize if it caused you trouble.

## Atari BASIC

Despite my protestation above, I actually did re-learn a bit of Atari BASIC 
for the post.  Please see [Atari docs](atari/README.md) for more details on what
it took to get the Atari 800 Emulator working.
