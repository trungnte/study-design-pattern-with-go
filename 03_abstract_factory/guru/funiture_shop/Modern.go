package main

import "fmt"

type ModernChair struct{}

func (a *ModernChair) hasLegs() {
	fmt.Println("Modern::Chair has LEDS")
}

func (a *ModernChair) sitOn() {
	fmt.Println("Modern::Chair sit ON")
}

type ModernSofa struct{}

func (a *ModernSofa) hasLegs() {
	fmt.Println("Modern::Sofa has LEDS")
}

func (a *ModernSofa) sitOn() {
	fmt.Println("Modern::Sofa sit ON")
}

type ModernCoffeTable struct{}

func (a *ModernCoffeTable) hasLegs() {
	fmt.Println("Modern::CoffeeTable has LEDS")
}
