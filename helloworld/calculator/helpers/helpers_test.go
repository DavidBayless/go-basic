package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "basics/helloworld/calculator/helpers"
)

var _ = Describe("Helpers", func() {
	Context("Add", func() {
		It("should return 0 when given 0 and 0", func() {
			sum := Add(0,0)
			Expect(sum).To(Equal(0))
		})

		It("should return 1 when given 1 and 0", func() {
			sum := Add(1,0)
			Expect(sum).To(Equal(1))
		})

		It("should return 1 when given 0 and 1", func() {
			sum := Add(0,1)
			Expect(sum).To(Equal(1))
		})

		It("should return -1 when given 0 and -1", func() {
			sum := Add(0,-1)
			Expect(sum).To(Equal(-1))
		})

		It("should return 2 when given 1 and 1", func() {
			sum := Add(1,1)
			Expect(sum).To(Equal(2))
		})

	})
	Context("subtract", func() {
		It("should return 0 when given 0 and 0", func() {
			difference := Subtract(0,0)
			Expect(difference).To(Equal(0))
		})

		It("should return 1 when given 1 and 0", func() {
			difference := Subtract(1,0)
			Expect(difference).To(Equal(1))
		})

		It("should return -1 when given 0 and 1", func() {
			difference := Subtract(0,1)
			Expect(difference).To(Equal(-1))
		})

		It("should return 1 when given 0 and -1", func() {
			difference := Subtract(0,-1)
			Expect(difference).To(Equal(1))
		})

		It("should return 0 when given 1 and 1", func() {
			difference := Subtract(1,1)
			Expect(difference).To(Equal(0))
		})

	})
	Context("multiply", func() {
		It("should return 0 when multiplied by 0", func() {
			Expect(Multiply(0, 78)).To(Equal(0))
			Expect(Multiply(1, 0)).To(Equal(0))
		})
		It("should return the two numbers multiplied together", func() {
			Expect(Multiply(1, 10)).To(Equal(10))
			Expect(Multiply(10, 10)).To(Equal(100))
			Expect(Multiply(-1, 10)).To(Equal(-10))
			Expect(Multiply(-10, -10)).To(Equal(100))
		})
	})
	Context("divide", func() {
		It("should return an error if divided by 0", func() {
			err, quotient := Divide(1, 0)
			Expect(err).To(Not(Equal(nil)))
			Expect(quotient).To(Equal(1.0))
		})
		It("should return 0 when the numerator is 0", func() {
			_, quotient := Divide(0, 10)
			Expect(quotient).To(Equal(0.0))
		})
		It("should return the correct quotient when divided", func() {
			_, quotient := Divide(10,3)
			Expect(quotient).To(Equal(10.0/3))
		})
	})
})
