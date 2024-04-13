package strategy

type EvictionAlgo interface {
	evict(c *Cache)
}
