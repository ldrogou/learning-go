package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"training.go/certgen/cert"
)

// PdfSaver pdf structure
type PdfSaver struct {
	OutputDir string
}

// New nouvelle instance
func New(outputDir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outputDir,
	}
	return p, nil
}

// Save implemente l'interface de sauvegarde de pdf
func (p *PdfSaver) Save(c cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(c.LabelTitle, false)
	pdf.AddPage()

	//Save file
	filename := fmt.Sprintf("%v.pdf", c.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to %v \n", path)

	return nil
}
