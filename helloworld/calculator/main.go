package main

import (
	"os"
	"fmt"
	"strconv"
	"basics/helloworld/calculator/helpers"
)

func main() {

	numStrOne := os.Args[1]
	operator := os.Args[2]
	numStrTwo := os.Args[3]

	numOne, _ := strconv.Atoi(numStrOne)
	numTwo, _ := strconv.Atoi(numStrTwo)

	if (operator == "+") {
		fmt.Println(helpers.Add(numOne, numTwo))
	} else if (operator == "-") {
		fmt.Println(helpers.Subtract(numOne, numTwo))
	} else if (operator == "x") {
		fmt.Println(helpers.Multiply(numOne, numTwo))
	} else if (operator == "/") {
		_, quotient := helpers.Divide(numOne, numTwo)
		fmt.Println(quotient)
	} else {
		fmt.Println("Incorrect Input!")
	}
}