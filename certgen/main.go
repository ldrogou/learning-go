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

	fileInput := flag.String("file", "", "fichier à traiter")
	outputType := flag.String("type", "pdf", "Output type.")
	flag.Parse()

	var certs []*cert.Cert
	var err error
	switch *fileInput {
	case "":
		fmt.Printf("Le nom du fichier doit être renseigné. fichier=%v", "")
		os.Exit(1)
	default:
		certs, err = parser.ParseCSV(*fileInput + ".csv")
		if err != nil {
			fmt.Printf("Erreur dans la lecture du fichier csv. err=%v \n", err)
			os.Exit(1)
		}
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
