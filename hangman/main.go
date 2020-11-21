package main

import (
	"fmt"
	"os"

	"training.go/hangman/dictionary"
	"training.go/hangman/hangman"
)

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Impossible de lire le dictionnaire de mots : %v \n", err)
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionary.PickWord())
	if err != nil {
		fmt.Printf("une erreur s'est produite : %v", err)
	}

	hangman.DrawWelcome()

	guess := ""

	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Lecture impossible à partir du terminal : %v .", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
