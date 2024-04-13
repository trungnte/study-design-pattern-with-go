package main

import "guru_strategy/strategy"

func main() {
	lfu := &strategy.Lfu{}
	cache := strategy.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	cache.ToString()

	lru := &strategy.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	cache.ToString()

	fifo := &strategy.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

	cache.ToString()

}
