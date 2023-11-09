package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type IHandler interface {
	handle(str string) string
	setNext(handler IHandler) IHandler
}

type Reciever struct {
	next IHandler
}

func (r *Reciever) setNext(next IHandler) IHandler {
	r.next = next
	return next
}

func (r *Reciever) handle(request string) string {
	return r.next.handle(request)
}

type Handler1 struct {
	next IHandler
}

func (h *Handler1) setNext(next IHandler) IHandler {
	h.next = next
	return next
}

func (h *Handler1) handle(request string) string {
	if request == "Hello, world!" {
		return "Hello!"
	}
	return h.next.handle(request)
}

type Handler2 struct {
	next IHandler
}

func (h *Handler2) setNext(next IHandler) IHandler {
	h.next = next
	return next
}

func (h *Handler2) handle(request string) string {
	if request == "Привет, мир!" {
		return "Привет!"
	}
	return ""
}

func main() {
	h2 := &Handler2{}
	h1 := &Handler1{}
	h1.setNext(h2)
	r := &Reciever{}
	r.setNext(h1)
	fmt.Println(r.handle("Hello, world!"))
	fmt.Println(r.handle("Привет, мир!"))
}
