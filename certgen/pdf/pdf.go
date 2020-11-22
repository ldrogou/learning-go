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

	// Background
	background(pdf)

	// --
	// header
	header(pdf, &c)

	//--
	// body
	pdf.SetFont("Arial", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelPresented, "C")
	pdf.Ln(30)

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, c.Name, "C")
	pdf.Ln(30)

	pdf.SetFont("Arial", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelParticipation, "C")
	pdf.Ln(30)

	pdf.SetFont("Arial", "I", 15)
	pdf.WriteAligned(0, 50, c.LabelDate, "C")

	//--
	// footer
	footer(pdf)

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

func background(pdf *gofpdf.Fpdf) {

	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("images/background.png", 0, 0,
		pageWidth, pageHeight,
		false, opts, 0, "")

}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {

	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "images/gopher.png"
	pdf.ImageOptions(filename, x+margin, 20,
		imageWidth, 0, false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename, x-margin, 20,
		imageWidth, 0, false, opts, 0, "")

	pdf.SetFont("Arial", "B", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
	pdf.Ln(30)
}

func footer(pdf *gofpdf.Fpdf) {

	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	imageWidth, imageHeigth := 45.0, 45.0
	pageWidth, pageHeigth := pdf.GetPageSize()
	filename := "images/stamp-certificate.png"
	x := pageWidth - imageWidth
	y := pageHeigth - imageHeigth
	pdf.ImageOptions(filename, x-margin, y-margin,
		imageWidth, 0, false, opts, 0, "")

}
