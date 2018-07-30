package oop_calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "basics/helloworld/oop_calculator"
	. "github.com/onsi/gomega"
)

var _ = Describe("OopCalculator", func() {
	Context("Calculator", func() {
		It("should tally the total", func() {
			calc := NewCalculator()
			calc.Add(10)
			Expect(calc.GetCurrentValue()).To(Equal(10.0))
			calc.Add(10)
			Expect(calc.GetCurrentValue()).To(Equal(20.0))
			calc.Subtract(5)
			Expect(calc.GetCurrentValue()).To(Equal(15.0))
			calc.Multiply(-2)
			Expect(calc.GetCurrentValue()).To(Equal(-30.0))
			calc.Divide(10)
			Expect(calc.GetCurrentValue()).To(Equal(-3.0))
			calc.Reset()
			Expect(calc.GetCurrentValue()).To(Equal(0.0))
		})
		It("should chain methods and tally the total", func() {
			calc := NewCalculator()
			calc.Add(10).Add(20).Subtract(-5).Multiply(0.5).Divide(2)
			Expect(calc.GetCurrentValue()).To(Equal(8.75))
		})
	})

})