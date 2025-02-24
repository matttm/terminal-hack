package utilities

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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
func BinarySearch(A []int, left, right, target int) int {
	return 0
}
func GetRandomRune() rune {
	// Define a range of Unicode code points for weird characters
	// These ranges include various non-Latin scripts, symbols, and other unusual characters
	ranges := []rune{'[', ']', '\\', '.', '/', '(', ')', '<', '>', '+'}

	// Select a random range
	r := ranges[rand.Intn(len(ranges))]

	return r
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
