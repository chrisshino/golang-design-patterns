package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

// --------

type PersonNameIterator struct {
	person *Person
	current int
}

func newPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current{
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2: 
		return p. person.LastName
	}
	panic("woops")
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for it := newPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
