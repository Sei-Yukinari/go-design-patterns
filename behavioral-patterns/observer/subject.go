// Package observer Subject
package observer

type Subject interface {
	register(Observer)
	deregister(Observer)
	notifyAll() []Observer
}
