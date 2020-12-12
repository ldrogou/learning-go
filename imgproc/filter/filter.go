package filter

import (
	"os"

	"github.com/disintegration/imaging"
)

type Filter interface {
	Process(srcPath, dstPath string) error
}

//Grayscale struct
type Grayscale struct{}

//Process implémentation de process
func (g Grayscale) Process(srcPath, dstPath string) error {

	// Open a test image.
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// Create a blurred version of the image.
	dst := imaging.Grayscale(src)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Save the resulting image as JPEG.
	opts := imaging.JPEGQuality(90)
	err = imaging.Save(dst, dstPath, opts)
	if err != nil {
		return err
	}

	return nil
}

//Blur struct
type Blur struct{}

//Process implémentation de process
func (g Blur) Process(srcPath, dstPath string) error {

	// Open a test image.
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// Create a blurred version of the image.
	dst := imaging.Blur(src, 3.5)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Save the resulting image as JPEG.
	opts := imaging.JPEGQuality(90)
	err = imaging.Save(dst, dstPath, opts)
	if err != nil {
		return err
	}

	return nil
}
