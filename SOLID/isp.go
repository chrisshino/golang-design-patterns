package main

type Document struct {

}

// better approach: split into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// combine interfaces

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combination + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}