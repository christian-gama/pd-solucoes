package providers

import (
	"strconv"
)

func Cpf() string {
	digits := generateRandomDigits(9)

	checkDigit1 := calculateCpfCheckDigit(digits)
	digits = append(digits, checkDigit1)
	checkDigit2 := calculateCpfCheckDigit(digits)
	digits = append(digits, checkDigit2)

	cpf := ""
	for _, digit := range digits {
		cpf += strconv.Itoa(digit)
	}

	return cpf
}

func calculateCpfCheckDigit(digits []int) int {
	const maxDigits = 11
	sum := 0

	for i := range digits {
		sum += digits[i] * (len(digits) + 1 - i)
	}
	checkDigit := sum % maxDigits

	if checkDigit < 2 {
		checkDigit = 0
	} else {
		checkDigit = maxDigits - checkDigit
	}

	return checkDigit
}
