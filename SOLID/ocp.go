package main

import "fmt"

type Color int

type Size int

const (
	red Color = iota
	green 
	blue
)

const (
	small Size = iota
	medium
	large
)


type Product struct {
	name string
	color Color
	size Size
}

type Filter struct {
	// 
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type BetterFilter struct {

}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product{
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

// Specification pattern

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	products := []Product{apple, tree}

	fmt.Printf("Green Products (old): \n")
	bf := BetterFilter{}

	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}
}