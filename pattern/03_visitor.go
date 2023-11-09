package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type shape interface {
	getType() string
	accept(Visitor)
}

type square struct {
	side int
}

type circle struct {
	radius int
}

type rectangle struct {
	l int
	b int
}

func (c *circle) getType() string {
	return "circle"
}

func (t *rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

type Visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

func (s *square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "square"
}

func (c *circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type middleCoordinates struct {
	x int
	y int
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {

	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForRectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}

func (a *middleCoordinates) visitForSquare(s *square) {

	fmt.Println("Calculating coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating coordinates for circle")
}
func (a *middleCoordinates) visitForRectangle(t *rectangle) {
	fmt.Println("Calculating coordinates for rectangle")
}

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
