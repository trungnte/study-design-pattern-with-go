package main

import "fmt"

type Container struct {
	Package
}

type Ship struct {
}

func (t *Ship) deliver(b *Container) (int, error) {
	fmt.Println("Ship:: deliver by sea in a CONTAINER")
	return 0, nil
}
