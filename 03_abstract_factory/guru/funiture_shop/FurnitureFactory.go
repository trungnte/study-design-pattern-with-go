package main

type FurnitureFactory interface {
	createChair() (Chair, error)
	createCoffeeTable() (CoffeeTable, error)
	createSofa() (Sofa, error)
}
