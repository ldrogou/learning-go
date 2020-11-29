package parser

import (
	"encoding/csv"
	"os"

	"training.go/certgen/cert"
)

//ParseCSV parse de fichier csv pour import
func ParseCSV(filename string) ([]*cert.Cert, error) {
	certs := make([]*cert.Cert, 0)

	f, err := os.Open(filename)
	if err != nil {
		return certs, err
	}
	defer f.Close()

	c := csv.NewReader(f)
	records, err := c.ReadAll()
	if err != nil {
		return certs, err
	}

	for _, rec := range records {
		c, err := cert.New(rec[0], rec[1], rec[2])
		if err != nil {
			return certs, err
		}
		certs = append(certs, c)
	}
	return certs, nil
}
