package calculator

import "github.com/google/wire"

type GeneralAdder interface {
	Add(x, y int) int
}

type Adder struct{}

func NewAdder() *Adder {
	return &Adder{}
}

func (a *Adder) Add(x, y int) int {
	return x + y
}

var AdderProvider = wire.NewSet(
	NewAdder,
	wire.Bind(new(GeneralAdder), new(*Adder)),
)
