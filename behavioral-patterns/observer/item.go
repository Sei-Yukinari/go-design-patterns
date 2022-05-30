// Package observer Concrete subject
package observer

import "fmt"

type item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}
func (i *item) updateAvailability() []Observer {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	return i.notifyAll()
}
func (i *item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() []Observer {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
	return i.observerList
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getEmail() == observer.getEmail() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
