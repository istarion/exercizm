package isbnverifier

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	digits := make([]int, 0)
	for _, r := range isbn {
		if r == ' ' || r == '-' {
			continue
		}
		if r >= '0' && r <= '9' {
			digits = append(digits, int(r-'0'))
		} else if len(digits) == 9 && unicode.ToUpper(r) == 'X' {
			digits = append(digits, 10)
		} else {
			return false
		}
		
	}
	if len(digits) != 10 {
		return false
	}
	sum := 0
	for i, d := range digits {
		sum += d * (10 - i)
	}
	return sum%11 == 0
}
