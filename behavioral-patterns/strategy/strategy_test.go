package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrategy(t *testing.T) {
	lfu := &lfu{}
	cache := initCache(lfu)
	t.Run("lfu", func(t *testing.T) {
		cache.add("a", "1")
		cache.add("b", "2")
		cache.add("c", "3")
		assert.IsType(t, cache.evictionAlgo, lfu)
	})
	t.Run("lru", func(t *testing.T) {
		lru := &lru{}
		cache.setEvictionAlgo(lru)
		cache.add("d", "4")
		assert.IsType(t, cache.evictionAlgo, lru)
	})
	t.Run("lru", func(t *testing.T) {
		fifo := &fifo{}
		cache.setEvictionAlgo(fifo)
		cache.add("e", "5")
		assert.IsType(t, cache.evictionAlgo, fifo)
	})
}
