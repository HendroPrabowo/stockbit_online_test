package main

import (
	"fmt"
	"sort"
)

func main() {
	arrWords := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(anagram(arrWords))
}

func anagram(arrWords []string) (answer [][]string) {
	listSortedWord := createListSortedWord(arrWords)
	result := map[string][]string{}
	for _, val := range arrWords {
		sortedWord := sortWord(val)
		if listSortedWord[sortedWord] {
			result[sortedWord] = append(result[sortedWord], val)
		}
	}
	for _, val := range result {
		answer = append(answer, val)
	}
	return
}

func createListSortedWord(arrWords []string) (list map[string]bool) {
	list = make(map[string]bool)
	for _, val := range arrWords {
		sortedWord := sortWord(val)
		list[sortedWord] = true
	}
	return
}

func isAnagram(word1, word2 string) bool {
	if sortWord(word1) == sortWord(word2) {
		return true
	}
	return false
}

func sortWord(word string) string {
	wordInRune := []rune(word)
	sort.Slice(wordInRune, func(i, j int) bool {
		return wordInRune[i] < wordInRune[j]
	})
	return string(wordInRune)
}
