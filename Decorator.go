package main

import "fmt"

type Coffee interface {
    Cost() int
    Description() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() int {
    return 5
}

func (c *SimpleCoffee) Description() string {
    return "джай кофе"
}

type CoffeeDecorator struct {
    coffee Coffee
}

func (c *CoffeeDecorator) Cost() int {
    return c.coffee.Cost()
}

func (c *CoffeeDecorator) Description() string {
    return c.coffee.Description()
}

type MilkDecorator struct {
    CoffeeDecorator
}

func (m *MilkDecorator) Cost() int {
    return m.coffee.Cost() + 2
}

func (m *MilkDecorator) Description() string {
    return m.coffee.Description() + ", молочко"
}

type SugarDecorator struct {
    CoffeeDecorator
}

func (s *SugarDecorator) Cost() int {
    return s.coffee.Cost() + 1
}

func (s *SugarDecorator) Description() string {
    return s.coffee.Description() + ", сахарок"
}

func main() {
    coffee := &SimpleCoffee{}
    fmt.Println(coffee.Description(), "стоит", coffee.Cost(), "тенге")

    milkCoffee := &MilkDecorator{coffee}
    fmt.Println(milkCoffee.Description(), "стоит", milkCoffee.Cost(), "рублей")

    sugarMilkCoffee := &SugarDecorator{milkCoffee}
    fmt.Println(sugarMilkCoffee.Description(), "стоит", sugarMilkCoffee.Cost(), "доллар")
}
