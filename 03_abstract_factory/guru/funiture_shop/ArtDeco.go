package main

import "fmt"

type ArtDecoChair struct{}

func (a *ArtDecoChair) hasLegs() {
	fmt.Println("ArtDeco::Chair has LEDS")
}

func (a *ArtDecoChair) sitOn() {
	fmt.Println("ArtDeco::Chair sit ON")
}

type ArtDecoSofa struct{}

func (a *ArtDecoSofa) hasLegs() {
	fmt.Println("ArtDeco::Sofa has LEDS")
}

func (a *ArtDecoSofa) sitOn() {
	fmt.Println("ArtDeco::Sofa sit ON")
}

type ArtDecoCoffeTable struct{}

func (a *ArtDecoCoffeTable) hasLegs() {
	fmt.Println("ArtDeco::CoffeeTable has LEDS")
}
