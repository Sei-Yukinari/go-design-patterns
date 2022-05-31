// Package strategy Concrete Strategy
package strategy

import "fmt"

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Printf("Evicting by %T strategy \n", c.evictionAlgo)
}
