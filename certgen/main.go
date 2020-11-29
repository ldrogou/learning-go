package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/certgen/cert"
	"training.go/certgen/html"
	"training.go/certgen/parser"
	"training.go/certgen/pdf"
)

func main() {

	outputType := flag.String("type", "pdf", "Output type.")
	flag.Parse()

	certs, err := parser.ParseCSV("input.csv")
	if err != nil {
		fmt.Printf("Erreur dans la lecture du fichier csv. err=%v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	switch *outputType {
	case "pdf":
		saver, err = pdf.New("output")
	case "html":
		saver, err = html.New("output")
	default:
		fmt.Printf("Unknown output type. got=%v\n", *outputType)
	}

	for _, cert := range certs {
		saver.Save(*cert)
	}
}
