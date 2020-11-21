package hangman

import "fmt"

// DrawWelcome message d'accueil
func DrawWelcome() {
	fmt.Println(
		`
 | |  | |   /\   | \ | |/ ____|  \/  |   /\   | \ | |
 | |__| |  /  \  |  \| | |  __| \  / |  /  \  |  \| |
 |  __  | / /\ \ |     | | |_ | |\/| | / /\ \ |     |
 | |  | |/ ____ \| |\  | |__| | |  | |/ ____ \| |\  |
 |_|  |_/_/    \_\_| \_|\_____|_|  |_/_/    \_\_| \_|
		`)
}

// Draw dessine le jeu
func Draw(g *Game, guess string) {
	drawTurns(g.Turnsleft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
     ____
    |    |
    |    o
    |   /|\
    |    | 
    |   / \
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 1:
		draw = `
     ____
    |    |
    |    o
    |   /|\
    |    | 
    |      
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 2:
		draw = `
     ____
    |    |
    |    o
    |   /|\
    |      
    |      
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 3:
		draw = `
     ____
    |    |
    |    o
    |      
    |      
    |      
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 4:
		draw = `
     ____
    |    |
    |     
    |       
    |      
    |      
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 5:
		draw = `
     __   
    |     
    |     
    |      
    |      
    |       
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 6:
		draw = `

    |
    |      
   _|_
  |   |_____
  |         |
  |_________|
                `
	case 7:
		draw = `
   ___
  |   |_____
  |         |
  |_________|
                `
	case 8:
		draw = `

		`
	}

	fmt.Println(draw)
}

func drawState(g *Game, guess string) {

	fmt.Print("Trouvées : ")
	drawLetters(g.FoundLetters)

	fmt.Print("Utilisées : ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess! \n")
	case "alreadyGuess":
		fmt.Printf("Lettre '%s' déjà utilisée ! \n", guess)
	case "badGuess":
		fmt.Printf("Mauvaise lettre, '%s' n'est pas dans le mot ! \n", guess)
	case "lost":
		fmt.Println("Tu as perdu :(! Le mot est ")
		drawLetters(g.Letters)
	case "won":
		fmt.Println("Tu as gagné :) ")
		drawLetters(g.Letters)

	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
