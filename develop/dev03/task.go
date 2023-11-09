package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	k := flag.Int("k", 3, "Select column")
	n := flag.Bool("n", false, "Sort by int")
	r := flag.Bool("r", false, "Sort reverse")
	u := flag.Bool("u", false, "Sort distinct")
	flag.Parse()
	data, err := readFile("text.txt")
	if err != nil {
		log.Fatalf("Cant open file, %v", err.Error())
	}
	switch {
	case *k > 0:
		data = sortByColumn(data, *k)
	case *n:
		data = sortByInt(data)
	case *r:
		data = sortReverse(data)
	case *u:
		data = sortDistinct(data)
	}
	if err = recordFile("result.txt", data); err != nil {
		log.Fatalf("Cant record file, %v", err.Error())
	}

}

func readFile(path string) ([]string, error) {
	result := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		result = append(result, str)
	}

	return result, nil
}

func recordFile(name string, data []string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range data {
		if _, err = file.WriteString(v + "\n"); err != nil {
			fmt.Printf("Cant write data %s, %v", v, err)
		}
	}

	return nil
}

func sortByColumn(strs []string, column int) []string {
	result := make([]string, 0)
	keys := make([]string, 0)
	m := make(map[string]string)
	column--

	for _, v := range strs {
		str := strings.Split(v, " ")
		if column >= len(str) {
			fmt.Println("Error! Column doesnt exist!")
			return nil
		}
		keys = append(keys, str[column])
		m[str[column]] = v
	}
	sort.Strings(keys)
	for _, v := range keys {
		result = append(result, m[v])
	}
	return result
}

func sortByInt(strs []string) []string {
	digits := make([]int, 0)
	nonDigits := make([]string, 0)
	result := make([]string, 0)
	m := make(map[string]string)
	for _, v := range strs {
		str := strings.Split(v, " ")
		m[str[0]] = v
		digit, err := strconv.Atoi(str[0])
		if err == nil {
			digits = append(digits, digit)
			continue
		}
		nonDigits = append(nonDigits, str[0])
	}
	sort.Ints(digits)
	for _, v := range digits {
		result = append(result, m[strconv.Itoa(v)])
	}

	for _, v := range nonDigits {
		result = append(result, m[v])
	}

	return result
}

func sortReverse(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}

func sortDistinct(s []string) []string {
	result := make([]string, 0)
	uniq := make(map[string]int)
	for i := range s {
		uniq[s[i]]++
	}
	for key := range uniq {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}
