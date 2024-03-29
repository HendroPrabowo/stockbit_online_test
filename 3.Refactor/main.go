package main

import (
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) <= 0 {
		return ""
	}
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}
	runes := []rune(str)
	wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}
	runes = []rune(wordsAfterFirstBracket)
	return string(runes[1:indexClosingBracketFound])
}
