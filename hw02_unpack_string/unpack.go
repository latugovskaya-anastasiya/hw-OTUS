package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ErrInvalidString error is returned if an invalid string is entered.
	ErrInvalidString = errors.New("invalid string")
	nilStr           string
)

// Unpack function unpacks the string with repeating characters.
func Unpack(s string) (string, error) {
	in := []rune(s)

	var (
		writeMeNextIter rune // we will write this variable in the next iteration
		allowInsertMe   bool // writeMeNextIter can be written in the following iteration
		afterSlash      bool // the next iteration will take into account that in this iteration there was "\"
		rtn             strings.Builder
	)

	inputRange := len(in)

	if inputRange == 0 {
		return nilStr, nil
	}

	for i := 0; i < inputRange; i++ {
		current := in[i]
		gotSlash := current == rune(92) // check the rune - "\"
		gotDigit := unicode.IsDigit(current)
		gotLetter := unicode.IsLetter(current)

		switch {
		case gotDigit && !afterSlash:
			if !allowInsertMe {
				return nilStr, ErrInvalidString
			}
			multiplier, _ := strconv.Atoi(string(current))
			rtn.WriteString(strings.Repeat(string(writeMeNextIter), multiplier))
			allowInsertMe = false

		case gotDigit && afterSlash:
			writeMeNextIter = current
			afterSlash = false
			allowInsertMe = true

		case gotSlash && !afterSlash:
			rtn.WriteString(string(writeMeNextIter))
			afterSlash = true
			allowInsertMe = false

		case gotSlash && afterSlash:
			writeMeNextIter = current
			afterSlash = false
			allowInsertMe = true

		case gotLetter:
			if !allowInsertMe {
				allowInsertMe = true
				writeMeNextIter = current
				continue
			}

			rtn.WriteString(string(writeMeNextIter))
			writeMeNextIter = current
			allowInsertMe = true

		case gotLetter && afterSlash:
			return nilStr, ErrInvalidString
		}
	}

	lastSymbol := in[inputRange-1]
	if allowInsertMe {
		rtn.WriteString(string(lastSymbol))
	}

	return rtn.String(), nil
}
