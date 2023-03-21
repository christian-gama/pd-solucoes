package providers

import (
	"strconv"
	"strings"
)

func Cnpj() string {
	baseCNPJ := generateRandomDigits(12)
	firstDigit := calculateCnpjCheckDigit(baseCNPJ, []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	secondDigit := calculateCnpjCheckDigit(
		append(baseCNPJ, firstDigit),
		[]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
	)

	baseCNPJ = append(baseCNPJ, firstDigit, secondDigit)
	return convertToString(baseCNPJ)
}

func calculateCnpjCheckDigit(digits []int, weightPattern []int) int {
	sum := 0
	for i, digit := range digits {
		sum += digit * weightPattern[i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func convertToString(cnpjDigits []int) string {
	strDigits := make([]string, len(cnpjDigits))
	for i, digit := range cnpjDigits {
		strDigits[i] = strconv.Itoa(digit)
	}

	return strings.Join(strDigits, "")
}
