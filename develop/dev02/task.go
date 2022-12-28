package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(s string) (string, error) {
	input := []rune(s)
	var result []rune

	if len(input) <= 1 {
		return s, nil
	}
	if unicode.IsDigit(input[0]) {
		return "", errors.New("input string must start with a letter")
	}

	for l, r := 0, 1; r < len(input); l, r = l+2, r+2 {
		lIsLetter, rIsLetter := unicode.IsLetter(input[l]), unicode.IsLetter(input[r])
		lIsDigit, rIsDigit := unicode.IsDigit(input[l]), unicode.IsDigit(input[r])

		if (lIsDigit && input[l] == '0') || (rIsDigit && input[r] == '0') {
			return "", errors.New("input string must not contain zeros")
		}

		switch {
		case lIsDigit && rIsLetter:
			if unicode.IsDigit(input[l-1]) {
				return "", errors.New("wrong input string")
			}

			if timesToAppend, err := strconv.Atoi(string(input[l])); err != nil {
				panic(err)
			} else {
				for i := 0; i < timesToAppend-1; i++ {
					result = append(result, input[l-1])
				}
				result = append(result, input[r])
			}
		case lIsLetter && rIsDigit:
			if timesToAppend, err := strconv.Atoi(string(input[r])); err != nil {
				panic(err)
			} else {
				for i := 0; i < timesToAppend; i++ {
					result = append(result, input[l])
				}
			}
		case lIsLetter && rIsLetter:
			result = append(result, input[l], input[r])
		case lIsDigit && rIsDigit:
			return "", errors.New("wrong input string")
		}
	}

	return string(result), nil
}

func main() {
	s := "a4bc2d5e"
	res, err := Unpack(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
