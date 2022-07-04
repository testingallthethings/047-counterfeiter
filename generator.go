package drivinglicence

import (
	"errors"
	"fmt"
)

//go:generate  counterfeiter -o mock . Applicant
type Applicant interface {
	IsOver17() bool
	HoldsLicence() bool
	GetInitials() string
	GetDOB() string
}

//go:generate  counterfeiter -o mock . Logger
type Logger interface {
	LogStuff(v string)
}

//go:generate  counterfeiter -o mock . RandomNumbersGenerator
type RandomNumbersGenerator interface {
	GetRandomNumbers(len int) string
}

type NumberGenerator struct {
	l Logger
	r RandomNumbersGenerator
}

func (g NumberGenerator) Generate(a Applicant) (string, error) {
	if a.HoldsLicence() {
		g.l.LogStuff("Duplicate Applicant, you can only hold one licence")
		return "", errors.New("Duplicate Applicant, you can only hold one licence")
	}

	if !a.IsOver17() {
		g.l.LogStuff("Underaged Applicant, you must be 17 for applicant licence")
		return "", errors.New("Underaged Applicant, you must be 17 for applicant licence")
	}

	n := fmt.Sprintf(
		"%s%s",
		a.GetInitials(),
		a.GetDOB(),
	)

	num := 16 - len(n)

	return fmt.Sprintf("%s%s", n, g.r.GetRandomNumbers(num)), nil
}

func NewNumberGenerator(l Logger, r RandomNumbersGenerator) NumberGenerator {
	return NumberGenerator{l, r}
}
