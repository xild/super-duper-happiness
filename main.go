package main

func Reverse(word string) (result string) {
	stringLen := len(word)
	runes := make([]rune, stringLen)
	for _, rune := range word {
		stringLen--
		runes[stringLen] = rune
	}
	return string(runes[stringLen:])
}
