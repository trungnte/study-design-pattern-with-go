package main

type FurnitureService struct {
	furniture   FurnitureFactory
	chair       Chair
	coffeeTable CoffeeTable
	sofa        Sofa
}

func (f *FurnitureService) GetFurniture(style string) (FurnitureFactory, error) {
	var fur FurnitureFactory
	if style == "ArtDeco" {
		fur = &ArtDecoFurnitureFactory{}
	} else if style == "Modern" {
		fur = &ModernFurnitureFactory{}
	} else {
		fur = &VictorianFurnitureFactory{}
	}
	return fur, nil
}

func (f *FurnitureService) CreateFurniture(style string) {
	f.furniture, _ = f.GetFurniture(style)
	f.chair, _ = f.furniture.createChair()
	f.coffeeTable, _ = f.furniture.createCoffeeTable()
	f.sofa, _ = f.furniture.createSofa()
}
