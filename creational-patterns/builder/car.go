package car

import "strconv"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor Color = "green"
	RedColor   Color = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels  Wheels = "steel"
)

type Car interface {
	Drive() error
	Stop() error
}

type car struct {
	topSpeed Speed
	wheel    Wheels
	color    Color
}

func (c *car) Drive() string {
	return "Driving at speed: " + strconv.FormatFloat(float64(c.topSpeed), 'f', -1, 64)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}

type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Car
}

type carBuilder struct {
	color Color
	wheel Wheels
	speed Speed
}

func New() *carBuilder {
	return &carBuilder{}
}

func (cb *carBuilder) Paint(color Color) *carBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Wheels(w Wheels) *carBuilder {
	cb.wheel = w
	return cb
}

func (cb *carBuilder) TopSpeed(s Speed) *carBuilder {
	cb.speed = s
	return cb
}

func (cb *carBuilder) Build() *car {
	return &car{
		color:    cb.color,
		wheel:    cb.wheel,
		topSpeed: cb.speed,
	}
}
