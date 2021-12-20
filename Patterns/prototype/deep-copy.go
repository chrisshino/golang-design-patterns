package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}


type Person struct {
	Name string
	Address *Address
}

func main() {
	john := Person{"John", &Address{"123 London", "London", "UK"}}	
	jane := john

	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker"

	fmt.Print(john.Address, jane.Address)
}