package main

import (
	"fmt"
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

func main() {
	str := []string{"пятак", "Пасечник", "тяпка", "пятка", "сяпка", "р", "песчаник", "Листок", "слиток", "песчинка", "столик"}

	result := anagram(str)
	fmt.Println(result)
	for i, v := range result {
		fmt.Printf("Слово: %s, анаграмы: %s\n", i, v)
	}

}

func anagram(strs []string) map[string][]string {
	result := make(map[string][]string)
	var key string
	for k := 0; k < len(strs); k++ {
		if strs[k] == "" {
			continue
		}
		key = strings.ToLower(strs[k])
		for i := 1; i < len(strs); i++ {
			lStr := strings.ToLower(strs[i])
			if key == lStr {
				continue
			}
			keyRune := []rune(key)
			strRune := []rune(lStr)
			if len(keyRune) != len(strRune) {
				continue
			}

			sort.Slice(keyRune, func(i, j int) bool {
				return keyRune[i] <= keyRune[j]
			})
			sort.Slice(strRune, func(i, j int) bool {
				return strRune[i] <= strRune[j]
			})
			if string(keyRune) != string(strRune) {
				continue
			}
			result[key] = append(result[key], lStr)
			strs[i] = ""
		}
		sort.Strings(result[key])
	}
	return result
}
