// Package observer Observer
package observer

type Observer interface {
	update(string)
	getEmail() string
}
