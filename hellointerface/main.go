package main

import (
	"fmt"
	"strconv"
)

type Artisan interface {
	Metier() string
}

type Boulanger struct {
	Debut int
}

type Peintre struct {
	Chantier int
}

func (b Boulanger) Metier() string {
	return "Dans mon métier je commence à " + strconv.Itoa(b.Debut)
}

func (p Peintre) Metier() string {
	return "Je possède " + strconv.Itoa(p.Chantier) + " chantiers "
}
func main() {
	artisans := []Artisan{Boulanger{4}, Peintre{20}}

	for _, artisan := range artisans {
		fmt.Println(artisan.Metier())
		fmt.Printf("type de l'artisan %T \n", artisan)

		determineArtisan(artisan)
		// if t, ok := artisan.(Boulanger); ok {
		// fmt.Printf("Je suis un Boulanger %v \n", t.Metier())
		// }
	}
}

func determineArtisan(a Artisan) {
	switch v := a.(type) {
	case Boulanger:
		fmt.Printf("Le boulanger se leve à %v heure \n", v.Debut)
	case Peintre:
		fmt.Printf("Le peintre à besoin de %v outils \n", v.Chantier)
	}
}
