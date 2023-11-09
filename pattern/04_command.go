package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type light struct {
	isRunning bool
}

func (l *light) on() {
	l.isRunning = true
	fmt.Println("Turning light on")
}

func (l *light) off() {
	l.isRunning = false
	fmt.Println("Turning light off")
}

func main() {
	light := &light{}
	onCommand := &onCommand{
		device: light,
	}
	offCommand := &offCommand{
		device: light,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
