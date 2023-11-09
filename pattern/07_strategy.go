package main

import (
	"fmt"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type sortStrategy interface {
	sort(digits []int)
}

type context struct {
	sortStrategy sortStrategy
}

func (c *context) setSortStrategy(s sortStrategy) {
	c.sortStrategy = s
}

func (c *context) getSorted(digits []int) {
	c.sortStrategy.sort(digits)
}

type bubbleSort struct {
}

func (b *bubbleSort) sort(digits []int) {
	fmt.Println("Using bubble sort to sort digits")
}

type quickSort struct {
}

func (q *quickSort) sort(digits []int) {
	fmt.Println("Using quick sort to sort digits")
}

type insertionSort struct {
}

func (i *insertionSort) sort(digits []int) {
	fmt.Println("Using insertion sort to sort digits")
}

func main() {
	digits := []int{15, 22, 11, -2, 98, 25}
	ctx := &context{}
	q := &quickSort{}
	b := &bubbleSort{}
	i := &insertionSort{}
	ctx.setSortStrategy(q)
	ctx.getSorted(digits)
	ctx.setSortStrategy(b)
	ctx.getSorted(digits)
	ctx.setSortStrategy(i)
	ctx.getSorted(digits)
}
