// Package utilities provides helper functions for word generation, random character creation,
// and hex offset formatting used in the terminal hacking game.
package utilities

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// GetWordList retrieves a list of words from the system dictionary file.
// It filters words by the specified length and returns a random selection.
// Returns count words, each with the specified length.
func GetWordList(count, length int) ([]string, error) {
	data, err := os.ReadFile("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}
	filter := func(array []string, cb func(string) bool) []string {
		ret := []string{}
		for _, w := range array {
			if cb(w) {
				ret = append(ret, w)
			}
		}
		return ret
	}
	s := string(data)
	words := strings.Split(s, "\n")
	words = filter(words, func(s string) bool { return len(s) == length })
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	// res, err := http.Get(fmt.Sprintf("https://random-word-api.herokuapp.com/word?number=%d&length=%d", count, length))
	// if err != nil {
	// 	panic(err)
	// }
	// _body, _ := io.ReadAll(res.Body)
	// var words []string
	// json.Unmarshal(_body, &words)
	return words[:count], nil
}

// GetRandomRune returns a random special character from a predefined set.
// Used for generating "dud" characters to fill spaces between words.
func GetRandomRune() rune {
	// Define a range of Unicode code points for weird characters
	// These ranges include various non-Latin scripts, symbols, and other unusual characters
	ranges := []rune{'[', ']', '\\', '.', '/', '(', ')', '<', '>', '+'}

	// Select a random range
	r := ranges[rand.Intn(len(ranges))]

	return r
}

// GenerateUnicodeString generates a single-character string using a random special character.
func GenerateUnicodeString() string {
	n := 1
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = GetRandomRune()
	}
	return string(runes)
}

// GenerateRandomStrings generates count random single-character strings.
// Used to fill the game grid with "dud" characters between words.
func GenerateRandomStrings(count int) []string {
	n := count
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = GenerateUnicodeString()
	}
	return s
}

// GenerateHexOffsets generates a list of hexadecimal memory addresses.
// Used to display memory offsets in the left column of the game grid.
// lpadding adds trailing spaces to each hex string for alignment.
func GenerateHexOffsets(count, lpadding int) []string {
	ans := []string{}
	var padSb strings.Builder
	for i := 0; i < lpadding; i++ {
		padSb.WriteRune(' ')
	}
	padding := padSb.String()
	for i := 0; i < count; i++ {
		hex := fmt.Sprintf("0x%04x%s", i*8, padding)
		ans = append(ans, hex)
	}
	return ans
}
