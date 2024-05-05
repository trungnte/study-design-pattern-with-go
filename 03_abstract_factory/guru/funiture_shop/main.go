package main

func main() {
	victorianFurniture := FurnitureService{}
	victorianFurniture.CreateFurniture("Victorian")

	modernFurniture := FurnitureService{}
	modernFurniture.CreateFurniture("Modern")

	artDecoFurniture := FurnitureService{}
	artDecoFurniture.CreateFurniture("ArtDeco")
}
