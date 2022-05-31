// Package state ConcreteStates
package state

type State string

const (
	Start      State = "start"
	Inprogress State = "inprogress"
	End        State = "end"
)

type StartState struct{}

func (s *StartState) handle(c *Context) {
	c.name = string(Start)
	c.currentState = &StartState{}
}

type InprogressState struct{}

func (s *InprogressState) handle(c *Context) {
	c.name = string(Inprogress)
	c.currentState = &InprogressState{}
}

type StopState struct{}

func (s *StopState) handle(c *Context) {
	c.name = string(End)
	c.currentState = &StopState{}
}
