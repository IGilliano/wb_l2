package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Smartphone struct {
	brand string
	OS    string
}

type ISmartphone interface {
	printInfo()
}

func (s *Smartphone) printInfo() {
	fmt.Printf("This is %s smartphone, working on %s", s.brand, s.OS)
}

type IPhone struct {
	smartphone Smartphone
}

func (i *IPhone) printInfo() {
	fmt.Printf("This is %s smartphone, working on %s\n", i.smartphone.brand, i.smartphone.OS)
}

func NewIphone() ISmartphone {
	return &IPhone{
		smartphone: Smartphone{
			brand: "Apple",
			OS:    "iOS",
		},
	}
}

type S22 struct {
	smartphone Smartphone
}

func (s *S22) printInfo() {
	fmt.Printf("This is %s smartphone, working on %s\n", s.smartphone.brand, s.smartphone.OS)
}

func NewS22() ISmartphone {
	return &S22{
		smartphone: Smartphone{
			brand: "Samsung",
			OS:    "Android",
		},
	}
}

func NewSmartphone(name string) ISmartphone {
	switch name {
	case "IPhone":
		return NewIphone()
	case "S22":
		return NewS22()
	default:
		fmt.Println("Error: wrong type of phone")
		return nil
	}
}

func main() {
	phone1 := NewSmartphone("IPhone")
	phone1.printInfo()
	phone2 := NewSmartphone("S22")
	phone2.printInfo()
}
