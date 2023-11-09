package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- `qwe\4\5` => qwe45 (*)
	- `qwe\45` => qwe44444 (*)
	- `qwe\\5` => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	testCases := []string{"a4bc2d5e"}
	for i := range testCases {
		result, err := unpack(testCases[i])
		if err != nil {
			fmt.Printf("Error! %v\n", err.Error())
		} else {
			fmt.Println("String:", result)
		}
	}

}

func unpack(str string) (string, error) {
	var result []byte
	for i, v := range str {
		if v < 49 || v > 57 {
			result = append(result, str[i])
			continue
		}

		if i == 0 {
			return "", errors.New("incorrect string")
		}

		digit, _ := strconv.Atoi(string(v))
		for j := 0; j < digit-1; j++ {
			result = append(result, str[i-1])
		}

	}

	return string(result), nil
}
