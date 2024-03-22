package calculator

type Calculator struct {
	adder GeneralAdder
}

func NewCalculator(adder GeneralAdder) Calculator {
	return Calculator{
		adder,
	}
}

func (c Calculator) Add(x, y int) int {
	return c.adder.Add(x, y)
}
