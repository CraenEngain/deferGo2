package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(input string) (string, error) {
	var result strings.Builder
	var escape, previousWasDigit bool
	var lastRune rune

	for i, r := range input {
		switch {
		case escape:
			result.WriteRune(r)
			escape = false
		case r == '\\':
			if previousWasDigit {
				return "", fmt.Errorf("недопустимая строка: за цифрой следует экранированный символ в позиции %d", i)
			}
			escape = true
		case unicode.IsDigit(r):
			if previousWasDigit {
				return "", fmt.Errorf("недопустимая строка: последовательные цифры в позиции %d", i)
			}
			if lastRune == 0 {
				return "", fmt.Errorf("недопустимая строка: начинается с цифры в позиции %d", i)
			}
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", fmt.Errorf("недопустимая цифра в позиции %d: %v", i, err)
			}
			result.WriteString(strings.Repeat(string(lastRune), count-1))
			previousWasDigit = true
		default:
			result.WriteRune(r)
			lastRune = r
			previousWasDigit = false
		}
	}

	if escape {
		return "", fmt.Errorf("недопустимая строка: заканчивается символом escape")
	}

	return result.String(), nil
}

func main() {

}
