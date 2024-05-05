package main

type Transport interface {
	Deliver(p *Package) (int, error)
}
