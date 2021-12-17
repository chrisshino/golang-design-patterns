package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome int
}

// functional

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural 

type EmployeeFactory struct {
	position string
	annualIncome int
}

func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, e.position, e.annualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}




func main() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer, manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}