package convert_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

func TestConvert(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "convert suite")
}

var _ = Describe("Convert", func() {
	Context("when string represents an unsigned int value", func() {
		It("converts a string to an unsigned int", func() {
			result, err := Convert("42")
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(uint64(42)))
		})
	})

	Context("when string represents a signed int value", func() {
		It("converts a string to a signed int", func() {
			result, err := Convert("-42")
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(int64(-42)))
		})
	})

	Context("when string represents a float value", func() {
		It("converts a string to a float", func() {
			result, err := Convert("3.1415")
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(3.1415))
		})
	})

	Context("when string represents a bool value", func() {
		It("converts a string to a bool", func() {
			result, err := Convert("true")
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(true))
		})
	})
})
