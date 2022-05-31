// Package state
package state

type state interface {
	handle(c *Context)
}
