// Copyright 2019 Enrico Foltran. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Convert numbers from base 10 integers to base X strings and back again.
package baseconv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	Base2Alphabet  = "01"
	Base16Alphabet = "0123456789ABCDEF"
	Base36Alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	Base56Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"
	Base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Base64Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
)

var (
	ErrAlphabetEmptyString = errors.New("Alphabet cannot be an empty string")
	ErrSignCharInBase      = errors.New("Sign character found in converter base alphabet")
)

type BaseConverter struct {
	alphabet      string
	signChar      string
	decimalDigits string
}

func New(alphabet, signChar string) (*BaseConverter, error) {
	if strings.TrimSpace(alphabet) == "" {
		return nil, ErrAlphabetEmptyString
	}

	if signChar == "" {
		signChar = "-"
	}

	for _, d := range alphabet {
		if string(d) == signChar {
			return nil, ErrSignCharInBase
		}
	}

	return &BaseConverter{
		alphabet:      alphabet,
		signChar:      signChar,
		decimalDigits: "0123456789",
	}, nil
}

func (bc *BaseConverter) String() string {
	return fmt.Sprintf("<BaseConverter: base%d (%s)>", len(bc.alphabet), bc.alphabet)
}

func (bc *BaseConverter) Encode(n int64) string {
	neg, value := bc.convert(strconv.FormatInt(n, 10), bc.decimalDigits, bc.alphabet, "-")
	if neg {
		return bc.signChar + value
	}

	return value
}

func (bc *BaseConverter) Decode(s string) (int64, error) {
	neg, value := bc.convert(s, bc.alphabet, bc.decimalDigits, bc.signChar)
	if neg {
		return strconv.ParseInt("-"+value, 10, 64)
	}

	return strconv.ParseInt(value, 10, 64)
}

func (bc *BaseConverter) convert(s, fromAlphabet, toAlphabet, signChar string) (bool, string) {
	neg := false
	if string(s[0]) == signChar {
		s = s[1:]
		neg = true
	}

	x := 0
	for _, digit := range s {
		x = x*len(fromAlphabet) + strings.Index(fromAlphabet, string(digit))
	}

	res := ""
	if x == 0 {
		res = string(toAlphabet[0])
	} else {
		for x > 0 {
			digit := x % len(toAlphabet)
			res = string(toAlphabet[digit]) + res
			x = int(x / len(toAlphabet))
		}
	}

	return neg, res
}

var (
	Base2, _  = New(Base2Alphabet, "-")
	Base16, _ = New(Base16Alphabet, "-")
	Base36, _ = New(Base36Alphabet, "-")
	Base56, _ = New(Base56Alphabet, "-")
	Base62, _ = New(Base62Alphabet, "-")
	Base64, _ = New(Base64Alphabet, "$")
)
