package main

import "fmt"

type Box struct {
	Package
}

type Truck struct {
}

func (t *Truck) deliver(b *Box) (int, error) {
	fmt.Println("Truck:: deliver by land in a BOX")
	return 0, nil
}
