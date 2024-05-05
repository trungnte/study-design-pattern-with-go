package main

import "fmt"

// The strategy interface declares operations common to all
// supported versions of some algorithm. The context uses this
// interface to call the algorithm defined by the concrete
// strategies.
type Strategy interface {
	Execute(a, b float32) float32
}

// Concrete strategies implement the algorithm while following
// the base strategy interface. The interface makes them
// interchangeable in the context.
type ConcreteStrategyAdd struct {
}

func (c *ConcreteStrategyAdd) Execute(a, b float32) float32 {
	fmt.Println("ConcreteStrategyAdd::Execute")
	return a + b
}

type ConcreteStrategySubtract struct {
}

func (c *ConcreteStrategySubtract) Execute(a, b float32) float32 {
	fmt.Println("ConcreteStrategySubtract::Execute")
	return a - b
}

type ConcreteStrategyMultiply struct {
}

func (c *ConcreteStrategyMultiply) Execute(a, b float32) float32 {
	fmt.Println("ConcreteStrategyMultiply::Execute")
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecuteStrategy(a, b float32) float32 {
	return c.strategy.Execute(a, b)
}

func main() {
	var c Context
	c.SetStrategy(&ConcreteStrategyAdd{})
	rs := c.ExecuteStrategy(1, 2)
	fmt.Printf("Result: %f\n", rs)

	c.SetStrategy(&ConcreteStrategySubtract{})
	rs = c.ExecuteStrategy(1, 2)
	fmt.Printf("Result: %f\n", rs)

	c.SetStrategy(&ConcreteStrategyMultiply{})
	rs = c.ExecuteStrategy(1, 2)
	fmt.Printf("Result: %f\n", rs)
}
