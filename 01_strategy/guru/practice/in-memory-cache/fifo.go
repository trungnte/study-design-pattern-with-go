package main

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}
