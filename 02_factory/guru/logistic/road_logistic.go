package main

import "fmt"

type RoadLogistic struct {
	truck Truck
}

func (r *RoadLogistic) Deliver(p *Package) (int, error) {
	box := &Box{
		Package: *p,
	}
	fmt.Println("RoadLogistic::Deliver: encap package to BOX")
	return r.truck.deliver(box)
}

func (r *RoadLogistic) createTransport() (Transport, error) {
	fmt.Println("RoadLogistic::createTransport with Truck")
	return r, nil
}
