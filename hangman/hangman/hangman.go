package hangman

import (
	"fmt"
	"strings"
)

// Game struvture du jeu
type Game struct {
	State        string   // Game state
	Letters      []string // Lettres dans le mots à trouver
	FoundLetters []string // Bonnes suposission
	UsedLetters  []string // lettres utilisées
	Turnsleft    int      // Nombre de tours effectué
}

// New nouveau tour
func New(turns int, word string) (*Game, error) {
	if len(word) < 2 {
		return nil, fmt.Errorf("le mot à deviner doit être plus grand (le mot fait %v lettre)", word)
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		Turnsleft:    turns,
	}
	return g, nil
}

// MakeAGuess faire une supposition
func (g *Game) MakeAGuess(guess string) {

	switch g.State {
	case "won", "lost":
		return
	}
	m := map[string]int{"test": 2}

	fmt.Printf("ee %v", len(m))
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuess"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.FoundLetters, g.Letters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.lostTurn(guess)
		if g.Turnsleft <= 0 {
			g.State = "lost"
		}
	}

}

func (g *Game) lostTurn(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	g.Turnsleft--

}
func hasWon(founds []string, letters []string) bool {
	for i, l := range letters {
		if l != founds[i] {
			return false
		}
	}
	return true
}

// RevealLetter reporter les lettres
func (g *Game) RevealLetter(guess string) {

	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = l
		}
	}

}
func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if guess == l {
			return true
		}
	}
	return false

}
