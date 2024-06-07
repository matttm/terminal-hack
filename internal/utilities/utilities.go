package utilities

import (
	"math/rand"
	"time"
)

func GetWordList(count int) ([]string, error) {
	words := []string{}
	for i := 0; i < count; i++ {
		words = append(words, "test")
	}
	return words, nil
}
func BinarySearch(A []int, left, right, target int) int {
	return 0
}
func GetRandomRune() rune {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define a range of Unicode code points for weird characters
	// These ranges include various non-Latin scripts, symbols, and other unusual characters
	ranges := []struct {
		low, high int
	}{
		{0x2200, 0x22FF}, // Mathematical Operators
		{0x2300, 0x23FF}, // Miscellaneous Technical
		{0x2A00, 0x2AFF}, // Supplemental Mathematical Operators
		{0x27C0, 0x27EF}, // Miscellaneous Mathematical Symbols-A
		{0x2980, 0x29FF}, // Miscellaneous Mathematical Symbols-B
		{0x2B00, 0x2BFF}, // Miscellaneous Symbols and Arrows
		{0x25A0, 0x25FF}, // Geometric Shapes
	}

	// Select a random range
	r := ranges[rand.Intn(len(ranges))]

	// Generate a random code point within the selected range
	codePoint := rand.Intn(r.high-r.low+1) + r.low

	return rune(codePoint)
}
func GenerateUnicodeString() string {
	n := 1
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = GetRandomRune()
	}
	return string(runes)
}
func GenerateRandomStrings(count int) []string {
	n := count
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = GenerateUnicodeString()
	}
	return s
}
