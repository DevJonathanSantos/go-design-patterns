package main

import "fmt"

// Product interface
type Product interface {
	GetName() string
}

// ConcreteProductA
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

// ConcreteProductB
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

// Factory function
func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	// Using the factory to create products
	productA := CreateProduct("A")
	productB := CreateProduct("B")

	fmt.Println(productA.GetName()) // Output: Product A
	fmt.Println(productB.GetName()) // Output: Product B
}
