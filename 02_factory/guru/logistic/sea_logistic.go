package main

import "fmt"

type SeaLogistic struct {
	ship Ship
}

func (r *SeaLogistic) Deliver(p *Package) (int, error) {
	container := &Container{
		Package: *p,
	}
	fmt.Println("SeaLogistic::Deliver: encap package to CONTAINER")
	return r.ship.deliver(container)
}

func (r *SeaLogistic) createTransport() (Transport, error) {
	fmt.Println("SeaLogistic::createTransport with Ship")
	return r, nil
}
