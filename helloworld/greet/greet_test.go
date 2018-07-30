package greet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "basics/helloworld/greet"
)

var _ = Describe("Greet", func() {
	It("should return HELLO WORLD! when given an empty string", func() {
		Expect(Greet("")).To(Equal("HELLO WORLD!"))
	})

	It("should return HELLO PANDA BEARS! when given PANDA BEARS", func() {
		Expect(Greet("PANDA BEARS")).To(Equal("HELLO PANDA BEARS!"))
	})

	It("should return HELLO UNICORN! when given unicorn", func() {
		Expect(Greet("unicorn")).To(Equal("HELLO UNICORN!"))
	})
})
