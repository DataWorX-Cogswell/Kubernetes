package adder_suite_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

func TestAdder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "adder Suite")
}

var _ = Describe("Adder", func() {

	Context("when summands are positive", func() {

		It("adds two numbers", func() {
			sum, err := Add(2, 3)
			Expect(err).NotTo(HaveOccurred())
			Expect(sum).To(Equal(5))
		})

	})

	Context("when summand is negative", func() {

		It("returns an err", func() {
			_, err := Add(-1, -1)
			Expect(err).To(HaveOccurred())
		})
	})

})

var _ = Describe("Divider", func() {

	Context("When both dividend and divisor are valid", func() {

		It("yields a valid return", func() {
			result, err := Divide(10, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(5))
		})
	})

	Context("When both dividend and divisor are valid", func() {

		It("yields a valid return", func() {
			result, err := Divide(8, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(4))
		})
	})

	Context("When both dividend is valid and divisor is invalid", func() {

		It("yields an error", func() {
			result, err := Divide(10, 0)
			Expect(err).To(HaveOccurred())
			Expect(result).To(Equal(0))
		})
	})

})
