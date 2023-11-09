package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var fields int
	var delimiter string
	var separated bool

	flag.IntVar(&fields, "f", 0, "Select columns")
	flag.StringVar(&delimiter, "d", "", "Select delimiter")
	flag.BoolVar(&separated, "s", false, "Only with delimiter")
	flag.Parse()

	if fields == 0 {
		log.Fatal("Error! Select columns")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		result, _ := cut(text, fields, delimiter, separated)
		fmt.Println(result)
	}
}

func cut(str string, f int, d string, s bool) (string, bool) {
	if s && !strings.Contains(str, d) {
		return "", false
	}
	splitedStr := strings.Split(str, d)
	if f <= len(splitedStr) {
		return splitedStr[f-1], true
	}
	return "", false
}
