package cpf

import (
	"errors"
	"fmt"
	"regexp"
)

type Cpf struct {
	value string
}

const (
	minCpfLength = 11
	maxCpfLength = 14
)

func New(value string) (Cpf, error) {
	if !isValid(value) {
		return Cpf{}, errors.New("invalid_cpf")
	}
	return Cpf{value: value}, nil
}

func (c Cpf) Value() string {
	return c.value
}

func isValid(value string) bool {
	cpf := removeMask(value)
	if !hasValidLength(cpf) {
		return false
	}
	if !hasValidDigits(cpf) {
		return false
	}

	d1 := calculateDigit(cpf, 10)
	d2 := calculateDigit(cpf, 11)

	cpfDigits := cpf[len(cpf)-2:]
	return cpfDigits == fmt.Sprintf("%d%d", d1, d2)
}

func hasValidLength(value string) bool {
	return len(value) >= minCpfLength && len(value) <= maxCpfLength
}

func hasValidDigits(value string) bool {
	for i := 1; i < len(value); i++ {
		if value[i] != value[0] {
			return true
		}
	}
	return false
}

func calculateDigit(value string, factor int) int {
	total := 0
	for _, digit := range value {
		if factor > 1 {
			total += int(digit-'0') * factor
			factor--
		}
	}
	rest := total % 11
	if rest < 2 {
		return 0
	} else {
		return 11 - rest
	}
}

func removeMask(value string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(value, "")
}
