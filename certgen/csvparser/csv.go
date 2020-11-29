package csvparser

import (
	"encoding/csv"
	"os"

	"training.go/certgen/cert"
)

//CsvParse parse file csv
type CsvParse struct {
	filename string
	file     *os.File
}

//New nouvelle instance de file de csv
func New(filename string) (*CsvParse, error) {
	cf := filename + ".csv"
	f, err := os.Open(cf)
	if err != nil {
		return nil, err
	}

	pFile := &CsvParse{
		filename: cf,
		file:     f,
	}
	return pFile, nil
}

//Parse parse de fichier csv pour import
func (p *CsvParse) Parse() ([]*cert.Cert, error) {
	certs := make([]*cert.Cert, 0)

	c := csv.NewReader(p.file)
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
