package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age int
}

type tiredPerson struct {
	name string
	age int
}

func (p *person) SayHello() {
	fmt.Printf("Hi my name is %s \n", p.name)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Hi my name is %s im really tired", p.name)
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main () {
	p := NewPerson("James", 101)
	p.SayHello()
}