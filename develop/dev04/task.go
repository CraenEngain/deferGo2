package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to find all sets of anagrams in a dictionary
func findAnagrams(words []string) map[string][]string {

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	// Сортируем слова, чтобы гарантировать уникальность в наборах
	sort.Strings(words)

	// Вспомогательная функция для получения отсортированного представления слова
	sortedWord := func(word string) string {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		return string(runes)
	}

	// Сопоставление отсортированных слов с исходными словами
	sortedMap := make(map[string][]string)

	for _, word := range words {
		sorted := sortedWord(word)
		sortedMap[sorted] = append(sortedMap[sorted], word)
	}

	// Создание map, используя первое встреченное слово в качестве ключа
	result := make(map[string][]string)
	for _, words := range sortedMap {
		if len(words) > 1 {
			result[words[0]] = words
		}
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "окт"}
	anagrams := findAnagrams(words)
	for key, set := range anagrams {
		fmt.Printf("%s: %v\n", key, set)
	}
}
