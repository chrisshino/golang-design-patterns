package main

type Person struct {
	name string
	age int
	eyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// validation
	}
	return &Person{name, age, 2}
}

func main() {
	
}