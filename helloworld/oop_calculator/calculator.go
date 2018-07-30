package oop_calculator

type calculator struct {
	value float64
}

func (c *calculator) Add(num float64) *calculator {
	c.value += num
	return c
}

func (c *calculator) Subtract(num float64) *calculator {
	c.value -= num
	return c
}

func (c *calculator) Multiply(num float64) *calculator {
	c.value *= num
	return c
}

func (c *calculator) Divide(num float64) *calculator {
	c.value /= num
	return c
}

func (c *calculator) Reset() *calculator {
	c.value = 0
	return c
}

func (c *calculator) GetCurrentValue() float64 {
	return c.value
}

func NewCalculator() calculator {
	return calculator{0}
}