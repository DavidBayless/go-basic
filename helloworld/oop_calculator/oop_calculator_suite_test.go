package oop_calculator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOopCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OopCalculator Suite")
}
