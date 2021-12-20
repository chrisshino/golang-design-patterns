package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetAddress, City, Country string
}


type Person struct {
	Name string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}

	_ = d.Decode(&result)

	return &result
}


func main() {
	john := Person{"John", &Address{"123 London", "London", "UK"}, []string{"Chris", "Matt"}}	

	jane := john.DeepCopy()
	jane.Address = &Address{"108 girdwood", "Barrie", "ON"}

	fmt.Println(john.Address, jane.Address)
}