package supermarket_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"tddpractice/internal/supermarket"
)

var _ = Describe("Supermarket", func() {
	It("calculates the cost of groceries", func() {
		register := supermarket.NewRegister()

		register.Scan("apple") // 3
		register.Scan("apple") // 3
		register.Scan("bananas") // 1
		register.Scan("bananas") // 1
		register.Scan("orange") // 2

		Expect(register.Total()).To(Equal(10))
	})

	It("calculates the cost of groceries when there are different groceries", func() {
		register := supermarket.NewRegister()

		register.Scan("apple") // 3
		register.Scan("apple") // 3
		register.Scan("orange") // 2
		register.Scan("pie") // 10

		Expect(register.Total()).To(Equal(18))
	})
})

