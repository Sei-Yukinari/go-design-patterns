package decorator

// Component
type pizza interface {
	getPrice() int
}

// ConcreteComponent
type vegetablePizza struct {
}

func (v *vegetablePizza) getPrice() int {
	return 15
}

// ConcreteDecorators
type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}

// ConcreteDecorators
type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
