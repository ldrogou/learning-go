package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("golang", "maelle", "2020-11-03")
	if err != nil {
		t.Errorf("Cert should be valid err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be valid reference got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid expected='GOLANG COURSE', actual=%v", c.Course)
	}

	if c.Name != "MAELLE" {
		t.Errorf("Cert Name is not valid expected='MAELLE', actual=%v", c.Name)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Maelle", "2020-11-04")
	if err == nil {
		t.Errorf("Cert should not be valid")
	}
}

func TestCourseTooLong(t *testing.T) {
	_, err := New("ededederfregtrhthdfvdcxsvgnhggbvcvdhndfgezrfezqezd", "Maelle", "2020-11-04")
	if err == nil {
		t.Errorf("Cert should not be valid course is too long")
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2020-11-04")
	if err == nil {
		t.Errorf("Cert should not be valid")
	}
}

func TestNameTooLong(t *testing.T) {
	_, err := New("Python", "Maellevbkdsfvfdvndfsnvlkdfnvlkjdsnvkdsnvkdsfnvksdfnv", "2020-11-04")
	if err == nil {
		t.Errorf("Cert should not be valid name is too long")
	}
}

func TestCertiBadDate(t *testing.T) {
	_, err := New("Python", "Laurent", "2020-15-04")
	if err == nil {
		t.Errorf("Cert should not be valid. err=%v", err)
	}
}

func TestCertGoodDate(t *testing.T) {
	c, err := New("Python", "Laurent", "2020-11-04")
	if err != nil {
		t.Errorf("Cert should be valid. err=%v", err)
	}
	if c.LabelDate != "Date 04/11/2020" {
		t.Errorf("Cert Label not weel formatted. expected='Date 04/11/2020', actual=%v", c.LabelDate)
	}
}
