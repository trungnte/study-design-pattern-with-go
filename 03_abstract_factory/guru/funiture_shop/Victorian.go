package main

import "fmt"

type VictorianChair struct{}

func (a *VictorianChair) hasLegs() {
	fmt.Println("Victorian::Chair has LEDS")
}

func (a *VictorianChair) sitOn() {
	fmt.Println("Victorian::Chair sit ON")
}

type VictorianSofa struct{}

func (a *VictorianSofa) hasLegs() {
	fmt.Println("Victorian::Sofa has LEDS")
}

func (a *VictorianSofa) sitOn() {
	fmt.Println("Victorian::Sofa sit ON")
}

type VictorianCoffeTable struct{}

func (a *VictorianCoffeTable) hasLegs() {
	fmt.Println("Victorian::CoffeeTable has LEDS")
}
