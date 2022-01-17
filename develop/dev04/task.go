package main

import (
	"fmt"
	"sort"
	"strings"
)

func toLower(words []string) []string {
	var res []string
	for _, val := range words {
		res = append(res, strings.ToLower(val))
	}
	return res
}

func deleteRepeat(words []string) []string {
	var res []string
	m := make(map[string]bool)

	for _, v := range words {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func searchForAnagrams(words []string) map[string][]string {
	// Создаем результирующую мапу
	mapH := make(map[string][]string)

	// Заранее переводим все символы в нижний регистр и удаляем повторения
	wordsLower := toLower(words)
	finWords := deleteRepeat(wordsLower)

	for _, val := range finWords {
		word := []rune(val)
		sort.Slice(word, func(i, j int) bool {
			return word[j] > word[i]
		})

		wordString := string(word)
		mapH[wordString] = append(mapH[wordString], val)
	}
	resMap := make(map[string][]string)
	for _, val := range mapH {
		if len(val) > 1 {

			resMap[val[0]] = val
			sort.Strings(val)
		}
	}
	return resMap
}

func main() {
	mas := []string{"пятка", "слиток", "пятак", "черешня", "тяпка", "столик"}
	res := searchForAnagrams(mas)
	fmt.Println(res)
}
