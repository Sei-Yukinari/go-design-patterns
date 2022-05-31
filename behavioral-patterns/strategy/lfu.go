// Package strategy Concrete Strategy
package strategy

import "fmt"

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Printf("Evicting by %T strategy \n", c.evictionAlgo)
}
