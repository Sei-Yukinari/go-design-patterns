package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObserver(t *testing.T) {
	item := newItem("aaa")

	o1 := NewCustomer("aaa@example.com")
	o2 := NewCustomer("bbb@example.com")

	item.register(o1)
	item.register(o2)

	assert.Equal(t, item.observerList, item.updateAvailability())

	item.deregister(o1)
	item.deregister(o2)

	assert.Equal(t, item.observerList, item.updateAvailability())
}
