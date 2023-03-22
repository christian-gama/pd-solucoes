package validators

import (
	"regexp"
)

type cpfValidator struct {
	v string
}

// Cpf returns true if the string is a valid CPF. It will validate the format,
// the blacklist and the check digits. A valid cpf contains numbers only.
func Cpf(v string) bool {
	validator := cpfValidator{v}
	return validator.validate()
}

func (c *cpfValidator) validate() bool {
	if c.v == "" {
		return true
	}

	return c.validateFormat() &&
		c.validateBlacklist() &&
		c.validate1stDigit() &&
		c.validate2ndDigit()
}

func (c *cpfValidator) validateFormat() bool {
	formatRegex := `^\d{11}$`
	r := regexp.MustCompile(formatRegex)

	return r.MatchString(c.v)
}

func (c *cpfValidator) validateBlacklist() bool {
	blacklist := map[string]bool{
		"00000000000": true,
		"11111111111": true,
		"22222222222": true,
		"33333333333": true,
		"44444444444": true,
		"55555555555": true,
		"66666666666": true,
		"77777777777": true,
		"88888888888": true,
		"99999999999": true,
		"12345678909": true,
	}

	return !blacklist[c.v]
}

func (c *cpfValidator) validate1stDigit() bool {
	// Calculate the sum of products
	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(c.v[i] - '0')
		product := digit * (10 - i)
		sum += product
	}

	// Calculate the first verification digit
	remainder := sum % 11
	digit1 := 0
	if remainder > 1 {
		digit1 = 11 - remainder
	}

	// Check if the calculated digit matches the actual digit
	return digit1 == int(c.v[9]-'0')
}

func (c *cpfValidator) validate2ndDigit() bool {
	// Calculate the sum of products
	sum := 0
	for i := 0; i < 10; i++ {
		digit := int(c.v[i] - '0')
		product := digit * (11 - i)
		sum += product
	}

	// Calculate the second verification digit
	remainder := sum % 11
	digit2 := 0
	if remainder > 1 {
		digit2 = 11 - remainder
	}

	// Check if the calculated digit matches the actual digit
	return digit2 == int(c.v[10]-'0')
}
