package main

import "fmt"

// Rect structure de type rectangle
type Rect struct {
	Largeur, Hauteur int
}

// Area calcul de l'aire d'un rectangle
func (r Rect) Area() int {
	return r.Hauteur * r.Largeur
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect => Largeur %d, Hauteur %d", r.Largeur, r.Hauteur)
}
func main() {
	r := Rect{2, 4}
	fmt.Printf("L'aire du rectangle est : %v \n", r.Area())
	fmt.Printf("main() %v \n", r)

}
