package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Test struct {
	Args     []string
	Expected []string
}

var testcases = []Test{
	{Args: []string{"hello 4 world", "laptop 6 pc", "abc 2 zbc", "laptop 6 pc"}},
	{Args: []string{"2 Some", "Test 11", "string 1", "just 4 testing"}},
	{Args: []string{"-1 Testing", "5 string", "4", "", "error"}},
}

func TestSortByColumn(t *testing.T) {
	testcases[0].Expected = []string{"laptop 6 pc", "laptop 6 pc", "hello 4 world", "abc 2 zbc"}
	testcases[1].Expected = []string{"string 1", "Test 11", "just 4 testing", "2 Some"}
	testcases[2].Expected = nil
	columns := []int{3, 2, 3}

	for i, test := range testcases {
		result := sortByColumn(test.Args, columns[i])
		require.Equal(t, result, test.Expected)
	}
}

func TestSortByInt(t *testing.T) {
	testcases[0].Expected = []string{"hello 4 world", "laptop 6 pc", "abc 2 zbc", "laptop 6 pc"}
	testcases[1].Expected = []string{"2 Some", "Test 11", "string 1", "just 4 testing"}
	testcases[2].Expected = []string{"-1 Testing", "4", "5 string", "", "error"}

	for _, test := range testcases {
		result := sortByInt(test.Args)
		require.Equal(t, result, test.Expected)
	}
}

func TestSortReverse(t *testing.T) {
	testcases[0].Expected = []string{"laptop 6 pc", "laptop 6 pc", "hello 4 world", "abc 2 zbc"}
	testcases[1].Expected = []string{"string 1", "just 4 testing", "Test 11", "2 Some"}
	testcases[2].Expected = []string{"error", "5 string", "4", "-1 Testing", ""}

	for _, test := range testcases {
		result := sortReverse(test.Args)
		require.Equal(t, result, test.Expected)
	}
}

func TestSortDistinct(t *testing.T) {
	testcases[0].Expected = []string{"abc 2 zbc", "hello 4 world", "laptop 6 pc"}
	testcases[1].Expected = []string{"2 Some", "Test 11", "just 4 testing", "string 1"}
	testcases[2].Expected = []string{"", "-1 Testing", "4", "5 string", "error"}

	for _, test := range testcases {
		result := sortDistinct(test.Args)
		require.Equal(t, result, test.Expected)
	}
}
