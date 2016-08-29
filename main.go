package main

import "unicode/utf8"

func Reverse(word string) (result string) {
	stringLen := utf8.RuneCountInString(word)
	runes := make([]rune, stringLen)
	for _, rune := range word {
		stringLen--
		runes[stringLen] = rune
	}
	return string(runes[stringLen:])
}
