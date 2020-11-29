package parser

import (
	"training.go/certgen/cert"
)

//Parser interface de parse file
type Parser interface {
	Parse() ([]*cert.Cert, error)
}
