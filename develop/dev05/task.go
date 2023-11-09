package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	After     int
	Before    int
	Context   int
	Count     bool
	Ignore    bool
	Invert    bool
	Fixed     bool
	LineNum   bool
	String    string
	Condition bool
	Matches   int
}

func NewFlags() *Flags {
	return &Flags{
		After:   0,
		Before:  0,
		Context: 0,
		Count:   false,
		Ignore:  false,
		Invert:  false,
		Fixed:   false,
		LineNum: false,
		String:  "",
		Matches: 0,
	}
}

func main() {
	flags := NewFlags()
	flag.IntVar(&flags.After, "A", 0, "Print N strings after coincidence")
	flag.IntVar(&flags.Before, "B", 0, "Print N strings before coincidence")
	flag.IntVar(&flags.Context, "C", 0, "Print N strings around coincidence")
	flag.BoolVar(&flags.Count, "c", false, "Print number of strings")
	flag.BoolVar(&flags.Ignore, "i", false, "Ignore register")
	flag.BoolVar(&flags.Invert, "v", false, "Invert from strings")
	flag.BoolVar(&flags.Fixed, "F", false, "Exect coincidence")
	flag.BoolVar(&flags.LineNum, "n", false, "Print string number")
	flag.Parse()
	if flags.Context > flags.After {
		flags.After = flags.Context
	}
	if flags.Context > flags.Before {
		flags.Before = flags.Context
	}
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Missing args. Input flags, string and file name")
	}

	strs := args[:len(args)-1]
	flags.String = strings.Join(strs, " ")

	lines, err := readFile(args[len(args)-1])
	if err != nil {
		log.Printf("Error:%v", err.Error())
	}

	search(lines, flags)

}

func readFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Cant close file: %v", err.Error())
		}
	}(file)

	data, err := io.ReadAll(file)
	lines := string(data)
	return strings.Split(lines, "\r\n"), nil
}

func search(lines []string, f *Flags) {
	result := make(map[int][]string)
	var condition bool

	for i, v := range lines {
		if f.Ignore {
			v = strings.ToLower(v)
			f.String = strings.ToLower(f.String)
		}
		if f.Fixed {
			condition = f.String == v
		} else {
			condition = strings.Contains(v, f.String)
		}
		if f.Invert {
			condition = !condition
		}
		if condition {
			f.Matches++
			left, right := f.Before, f.After
			if left > i {
				left = i
			}

			if right > len(lines)-i-1 {
				right = len(lines) - i - 1
			}

			result[i] = append(result[i], lines[i-left:i+right+1]...)
		}
	}
	for _, slice := range result {
		for _, v := range slice {
			fmt.Println(v)
		}
	}
}
