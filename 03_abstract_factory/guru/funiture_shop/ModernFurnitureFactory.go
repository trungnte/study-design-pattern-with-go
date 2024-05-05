package main

import "fmt"

type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) createChair() (Chair, error) {
	fmt.Println("ModernFurnitureFactory:: create Chair")
	return &ModernChair{}, nil
}

func (m *ModernFurnitureFactory) createSofa() (Sofa, error) {
	fmt.Println("ModernFurnitureFactory:: create Sofa")
	return &ModernSofa{}, nil
}

func (m *ModernFurnitureFactory) createCoffeeTable() (CoffeeTable, error) {
	fmt.Println("ModernFurnitureFactory:: create CoffeeTable")
	return &ModernCoffeTable{}, nil
}
