package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestDivide(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Divide Suite")
}

var _ = Describe("Divide", func() {
	It("divides two numbers", func() {
		Expect(Divide(4, 2).value).To(Equal(float32(2)))
	})

	Context("when denominator is 0", func() {
		It("returns an error", func() {
			Expect(Divide(4, 0).err.Error()).To(Equal("division by 0"))
		})
	})
})
