package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://mon-expert-en-gestion.fr")
	if err != nil {
		fmt.Printf("Erreur dans la requete http : %v", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	file, err := os.Create("fichier.txt")
	defer file.Close()

	io.Copy(file, res.Body)
}
