package main

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type PC struct {
	Motherboard string
	CPU         string
	GraphicCard string
	RAM         string
	PowerSupply string
}

type Builder struct {
	pc PC
}
type IBuilder interface {
	setMotherboard()
	setCPU()
	setGraphicCard()
	setRAM()
	setPowerSupply()
	Build() *PC
}

func (b *Builder) setMotherboard() {
	b.pc.Motherboard = "ASUS ROG STRIX B660-F GAMING WIFI"
}

func (b *Builder) setCPU() {
	b.pc.CPU = " AMD Ryzen 9 5950X"
}

func (b *Builder) setGraphicCard() {
	b.pc.GraphicCard = " NVIDIA GeForce RTX 4070TI"
}

func (b *Builder) setRAM() {
	b.pc.RAM = "Kingston Fury Beast KF432C16BB1K2/32 DDR4"
}

func (b *Builder) setPowerSupply() {
	b.pc.PowerSupply = " DeepCool PF800"
}

func (b *Builder) Build() *PC {
	return &PC{
		Motherboard: b.pc.Motherboard,
		CPU:         b.pc.CPU,
		GraphicCard: b.pc.GraphicCard,
		RAM:         b.pc.RAM,
		PowerSupply: b.pc.PowerSupply,
	}
}
