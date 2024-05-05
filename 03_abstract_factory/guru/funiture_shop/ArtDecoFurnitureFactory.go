package main

import "fmt"

type ArtDecoFurnitureFactory struct{}

func (a *ArtDecoFurnitureFactory) createChair() (Chair, error) {
	fmt.Println("ArtDecoFurnitureFactory:: create Chair")
	return &ArtDecoChair{}, nil
}

func (a *ArtDecoFurnitureFactory) createSofa() (Sofa, error) {
	fmt.Println("ArtDecoFurnitureFactory:: create Sofa")
	return &ArtDecoSofa{}, nil
}

func (a *ArtDecoFurnitureFactory) createCoffeeTable() (CoffeeTable, error) {
	fmt.Println("ArtDecoFurnitureFactory:: create CoffeeTable")
	return &ArtDecoCoffeTable{}, nil
}
