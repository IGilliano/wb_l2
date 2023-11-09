package main

import (
	"fmt"
	"testing"
)

type Flag struct {
	field     int
	delimeter string
	separated bool
}

type Test struct {
	text     string
	flag     Flag
	expected string
}

var TestCases = []Test{
	{text: "This is a test string", flag: Flag{field: 2, delimeter: " ", separated: false}, expected: "is"},
	{text: "This.is.a.test.string", flag: Flag{field: 3, delimeter: ".", separated: false}, expected: "a"},
	{text: "Thisisateststring", flag: Flag{field: 4, delimeter: " ", separated: true}, expected: ""},
	{text: "This is a test string", flag: Flag{field: 4, delimeter: " ", separated: true}, expected: "test"},
}

func TestCut(t *testing.T) {
	for _, test := range TestCases {
		if result, _ := cut(test.text, test.flag.field, test.flag.delimeter, test.flag.separated); result != test.expected {
			fmt.Printf(result)
			t.Errorf("Result %s not equal to expected %s", result, test.expected)
		}
	}
}
