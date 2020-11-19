package cert

import (
	"fmt"
	"strings"
	"time"
)

var maxLenCourse int = 20

// Cert struct des donnéescertificat
type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type saver interface {
	save(c Cert) error
}

//New construit un nouveau certificat
func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course: c,
		Name:   n,

		LabelTitle:         fmt.Sprintf(" %v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This Certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date %v", d.Format("02/01/2006")),
	}
	return cert, nil
}

func validateCourse(course string) (string, error) {
	suf := " course"
	course, err := validEmptyString(course)
	if err != nil {
		return course, err
	}
	if !strings.HasSuffix(course, suf) {
		course = course + suf
	}
	return strings.ToTitle(course), nil
}

func validEmptyString(s string) (string, error) {
	c := strings.TrimSpace(s)
	if len(c) == 0 || len(c) > maxLenCourse {
		return c, fmt.Errorf("invalid course got=%v, len=%v", c, len(c))
	}
	return c, nil
}

func validateName(name string) (string, error) {
	name, err := validEmptyString(name)
	if err != nil {
		return name, err
	}
	return strings.ToTitle(name), nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
