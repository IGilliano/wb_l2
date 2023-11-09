package main

import (
	"fmt"
	"log"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Meal struct {
	Name   string
	ID     int
	Recipe Recipe
}

type Recipe struct {
	Ingridients string
	Time        time.Duration
}

type Menu struct {
	Meals []Meal
}

func NewMenu() *Menu {
	meal := Meal{ID: 1, Name: "fried eggs", Recipe: Recipe{"2 eggs and 2 sausages", 10 * time.Second}}
	return &Menu{Meals: []Meal{meal}}
}

type Cook struct {
}

func (c *Cook) CookMeal(r Recipe) *Meal {
	time.Sleep(r.Time)
	return &Meal{}
}

func NewCook() *Cook {
	return &Cook{}
}

type Order struct {
	mealId   int
	clientId int
}

// Фасад:
type Waiter struct {
	order Order
	menu  Menu
	cook  Cook
}

func NewWaiter(menu Menu, cook Cook) *Waiter {
	return &Waiter{menu: menu, cook: cook}
}

func (w *Waiter) GetOrder(clientId int) {
	var mealId int
	log.Println("What do you want to order?")
	_, err := fmt.Scanf("%d", &mealId)
	if err != nil {
		log.Printf("Cant take order: %v", err.Error())
	}

	w.order = Order{clientId: clientId, mealId: mealId}
}

func (w *Waiter) PlaceOrder() *Meal {
	var meal Meal
	for _, v := range w.menu.Meals {
		if v.ID == w.order.mealId {
			meal = v
		}
	}
	order := w.cook.CookMeal(meal.Recipe)
	order.Name = meal.Name
	return order
}

func main() {
	clientId := 1
	menu := NewMenu()
	cook := NewCook()
	waiter := NewWaiter(*menu, *cook)
	waiter.GetOrder(clientId)
	meal := waiter.PlaceOrder()
	fmt.Printf("Your order, %s, is ready, sir! Bon appetit\n", meal.Name)
}
