package utilities

func GetWordList(count int) ([]string, error) {
	words := []string{}
	for i := 0; i < count; i++ {
		words = append(words, "test")
	}
	return words, nil
}
