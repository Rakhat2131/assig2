package main

import "fmt"

type Product interface {
    Name() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Name() string {
    return "Product A"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Name() string {
    return "Product B"
}

type Factory interface {
    CreateProduct() Product
}

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
    return &ConcreteProductA{}
}

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
    return &ConcreteProductB{}
}

func main() {
    factoryA := &ConcreteFactoryA{}
    productA := factoryA.CreateProduct()
    fmt.Println("Product created by factory A:", productA.Name())

    factoryB := &ConcreteFactoryB{}
    productB := factoryB.CreateProduct()
    fmt.Println("Product created by factory B:", productB.Name())
}
