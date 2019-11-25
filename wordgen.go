package wordgen

import (
	"math/rand"
	"time"
)

func coin() bool {
	return rand.Float64() > 0.5
}

// New generates a new word with letters, numbers and symbols
func New(size int) string {
	str := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		str += string(rune(rand.Intn(93) + 33))
	}
	return str
}

// OnlyLetters generates a new word with only letters
func OnlyLetters(size int) string {
	rand.Seed(time.Now().UnixNano())
	str := ""
	high := 0
	low := 0
	for i := 0; i < size; i++ {
		if coin() {
			low = 65
			high = 90
		} else {
			low = 97
			high = 122
		}

		str += string(rune(rand.Intn(high-low) + low))
	}
	return str
}

// OnlyNumbers generates a new word with only numbers
func OnlyNumbers(size int) string {
	rand.Seed(time.Now().UnixNano())
	str := ""
	high := 57
	low := 48
	for i := 0; i < size; i++ {
		str += string(rune(rand.Intn(high-low) + low))
	}
	return str
}

// NotSymbols generates a new word with letters and numbers but not symbols
func NotSymbols(size int) string {
	str := ""
	for i := 0; i < size; i++ {
		if coin() {
			str += OnlyLetters(1)
			continue
		}
		str += OnlyNumbers(1)

	}
	return str
}
