package main

import "fmt"

type VictorianFurnitureFactory struct{}

func (v *VictorianFurnitureFactory) createChair() (Chair, error) {
	fmt.Println("VictorianFurnitureFactory:: create Chair")
	return &VictorianChair{}, nil
}

func (v *VictorianFurnitureFactory) createSofa() (Sofa, error) {
	fmt.Println("VictorianFurnitureFactory:: create Sofa")
	return &VictorianSofa{}, nil
}

func (v *VictorianFurnitureFactory) createCoffeeTable() (CoffeeTable, error) {
	fmt.Println("VictorianFurnitureFactory:: create CoffeeTable")
	return &VictorianCoffeTable{}, nil
}
