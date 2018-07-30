package helpers

import "errors"

func Add(numOne int, numTwo int) int {
	return numOne + numTwo
}

func Subtract(numOne int, numTwo int) int {
	return numOne - numTwo
}

func Multiply(numOne int, numTwo int) int {
	return numOne * numTwo
}

func Divide(numOne int, numTwo int) (error, float64) {
	if (numTwo == 0) {
		return errors.New("Cannot divide by 0"), float64(numOne)
	}

	return nil, float64(numOne) / float64(numTwo)

}