// Package observer Concrete observer
package observer

import "fmt"

type customer struct {
	email string
}

func NewCustomer(email string) *customer {
	return &customer{
		email: email,
	}
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.email, itemName)
}

func (c *customer) getEmail() string {
	return c.email
}
