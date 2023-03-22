package validators

import (
	"fmt"
	"strconv"
)

type cnpjValidator struct {
	v string
}

// Cnpj returns true if the string is a valid CNPJ. It will validate the format,
// and the check digits. A valid CNPJ contains numbers only.
func Cnpj(v string) bool {
	validator := cnpjValidator{v}

	return validator.validate()
}

func (c cnpjValidator) calculateDigit(cnpj string, pos int) int {
	weights := [][]int{
		{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
		{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
	}
	sum := 0
	for i, v := range cnpj[:pos] {
		digit, _ := strconv.Atoi(string(v))
		sum += digit * weights[pos-12][i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func (c *cnpjValidator) validate() bool {
	if len(c.v) != 14 {
		return false
	}

	firstDigit := c.calculateDigit(c.v, 12)
	secondDigit := c.calculateDigit(c.v+"0", 13)

	checkDigits := c.v[12:]
	return checkDigits == fmt.Sprintf("%d%d", firstDigit, secondDigit)
}
