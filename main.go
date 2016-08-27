package main

func Reverse(s string) (result string) {
	stringLen := len(s)
	runes := make([]rune, stringLen)
	for _, rune := range s {
		stringLen--
		runes[stringLen] = rune
	}
	return string(runes[stringLen:])
}
