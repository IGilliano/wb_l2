package main

import (
	"fmt"
	"testing"
)

type Test struct {
	args     string
	expected string
}

var TestCases = []Test{
	{args: "a4bc2d5e", expected: "aaaabccddddde"},
	{args: "abcd", expected: "abcd"},
	{args: "45", expected: ""},
	{args: "", expected: ""},
}

func TestUnpack(t *testing.T) {
	for _, test := range TestCases {
		if result, _ := unpack(test.args); result != test.expected {
			fmt.Printf(result)
			t.Errorf("Result %s not equal to expected %s", test.args, test.expected)
		}
	}
}
