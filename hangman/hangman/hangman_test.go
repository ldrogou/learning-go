package hangman

import (
	"fmt"
	"testing"
)

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contient la lettre %s. on obtient=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s ne contient pas la lettre %s. on obtient=%v", word, guess, hasLetter)
	}
}

func TestNewGame(t *testing.T) {
	word := ""
	_, err := New(3, word)
	if err == nil {
		t.Errorf("La mot à deviner doit faire plus de 2 lettre, il en fait %v", len(word))
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(5, "orvault")
	g.MakeAGuess("a")
	validState(t, "goodGuess", g.State)
}
func TestGameBadGuess(t *testing.T) {
	g, _ := New(5, "orvault")
	g.MakeAGuess("p")
	validState(t, "badGuess", g.State)
}
func TestGameAlreadyGuess(t *testing.T) {
	g, _ := New(5, "orvault")
	g.MakeAGuess("a")
	g.MakeAGuess("a")
	validState(t, "alreadyGuess", g.State)
}
func TestGameWon(t *testing.T) {
	g, _ := New(5, "orvault")
	g.MakeAGuess("o")
	g.MakeAGuess("r")
	g.MakeAGuess("v")
	g.MakeAGuess("a")
	g.MakeAGuess("u")
	g.MakeAGuess("t")
	g.MakeAGuess("l")
	g.MakeAGuess("l")
	fmt.Print(g)
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(5, "orvault")
	g.Turnsleft = 1
	g.MakeAGuess("Z")
	validState(t, "lost", g.State)
}

func validState(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("le state attendu est %v mais le restultat est %v", expected, actual)
	}
}
