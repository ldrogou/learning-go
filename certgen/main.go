package main

import (
	"fmt"
	"os"

	"training.go/certgen/cert"
	"training.go/certgen/pdf"
)

func main() {

	c, err := cert.New("Golang Prog", "User Test", "2020-11-21")
	if err != nil {
		fmt.Printf("Erreur dans la création d'un CERT. err=%v", err)
		os.Exit(1)
	}
	var saver cert.Saver
	saver, err = pdf.New("output")
	saver.Save(*c)
}
