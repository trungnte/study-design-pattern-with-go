package strategy

import "fmt"

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	fmt.Printf("InitCache was called storage map: %v\n", storage)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func (c *Cache) ToString() {
	for k, v := range c.storage {
		fmt.Printf("c.storage[%v] = %v\n", k, v)
	}
	fmt.Printf("capacity: %d, maxCapacity: %d\n", c.capacity, c.maxCapacity)
	fmt.Println("---------------------------------------------\n")
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
