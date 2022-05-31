// Package strategy Concrete Strategy
package strategy

import "fmt"

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Printf("Evicting by %T strategy \n", c.evictionAlgo)
}
