package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Test struct {
	args     []string
	expected map[string][]string
}

var TestCases = []Test{
	{args: []string{"пятак", "Пасечник", "тяпка", "пятка", "сяпка", "р", "песчаник", "Листок", "слиток", "песчинка", "столик"},
		expected: map[string][]string{"листок": {"слиток", "столик"}, "пасечник": {"песчаник", "песчинка"}, "пятак": {"пятка", "тяпка"}}},
	{args: []string{"worker", "Привет", "23", "мир", "сяпка", "14", "Нет", "Дубликатов"},
		expected: map[string][]string{}},
	{args: []string{""},
		expected: map[string][]string{}},
}

func TestAnagram(t *testing.T) {
	for _, test := range TestCases {
		if result := anagram(test.args); !reflect.DeepEqual(result, test.expected) {
			fmt.Println(result)
			t.Errorf("Result %s not equal to expected %s", test.args, test.expected)
		}
	}
}
