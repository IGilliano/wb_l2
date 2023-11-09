package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type state interface {
	UpdateOrder(i string)
	setDeliveryMethod(d string)
	trackOrder()
}

type order struct {
	items          string
	deliveryMethod string
	placed         state
	shipped        state
	arrived        state
	current        state
}

func newOrder(items string) *order {
	o := &order{items: items}

	placed := &orderPlacedState{order: o}
	shipped := &orderShippedState{order: o}
	arrived := &orderArrivedState{order: o}
	o.placed = placed
	o.shipped = shipped
	o.arrived = arrived

	o.current = placed
	return o
}

func (o *order) setState(s state) {
	o.current = s
}

type orderPlacedState struct {
	order *order
}

func (p *orderPlacedState) UpdateOrder(i string) {
	p.order.items = i
}

func (p *orderPlacedState) setDeliveryMethod(d string) {
	p.order.deliveryMethod = d
	p.order.setState(p.order.shipped)
}

func (p *orderPlacedState) trackOrder() {
	log.Println("Order is not shipped yet")
}

type orderShippedState struct {
	order *order
}

func (p *orderShippedState) UpdateOrder(i string) {
	log.Println("Cant change order: order is already shipped")
}

func (p *orderShippedState) setDeliveryMethod(d string) {
	log.Println("Cant change delivery method: order is already shipped")
}

func (p *orderShippedState) trackOrder() {
	log.Println("Order is moving...")
}

type orderArrivedState struct {
	order *order
}

func (a *orderArrivedState) UpdateOrder(i string) {
	log.Println("Cant change order: order is already arrived")
}

func (a *orderArrivedState) setDeliveryMethod(d string) {
	log.Println("Cant change delivery method: order is already arrived")
}

func (a *orderArrivedState) trackOrder() {
	log.Println("Order arrived")
}

func main() {
	o := newOrder("Food")
	o.current.UpdateOrder("Drinks")
	o.current.setDeliveryMethod("Courier")
	o.current.trackOrder()
	o.setState(o.shipped)
	fmt.Println("Order is shipped")
	o.current.UpdateOrder("Food")
	o.current.setDeliveryMethod("Pickup")
	o.current.trackOrder()
	o.setState(o.arrived)
	fmt.Println("Order is arrived")
	o.current.UpdateOrder("Drinks")
	o.current.setDeliveryMethod("Courier")
	o.current.trackOrder()
}
