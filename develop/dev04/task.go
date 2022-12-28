package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func GetAnagrams(words *[]string) *map[string][]string {
	allSortedWords := make(map[string]string)
	for _, word := range *words {
		lowerWord := strings.ToLower(word)
		wordSorted := sortString(lowerWord)
		found := false
		for _, v := range allSortedWords {
			if v == wordSorted {
				found = true
				break
			}
		}
		if !found {
			allSortedWords[lowerWord] = wordSorted
		}
	}

	result := make(map[string][]string, len(allSortedWords))
	for _, word := range *words {
		lowerWord := strings.ToLower(word)
		wordSorted := sortString(lowerWord)
		for k, v := range allSortedWords {
			if wordSorted == v {
				result[k] = append(result[k], lowerWord)
			}
		}
	}

	for k, v := range result {
		if len(v) == 1 {
			delete(result, k)
		} else {
			sort.Strings(result[k])
		}
	}

	return &result
}

func sortString(word string) string {
	wordRunes := []rune(word)
	sort.Slice(wordRunes, func(i, j int) bool {
		return wordRunes[i] < wordRunes[j]
	})

	wordSorted := string(wordRunes)
	return wordSorted
}
