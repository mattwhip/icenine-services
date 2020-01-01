package uuid

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UUID tests", func() {

	Describe("Creating a new UUID", func() {
		It("should succeed", func() {
			By("Creating UUID")
			u, err := New()
			Expect(err).To(BeNil())
			Expect(len(u)).To(BeNumerically(">", 0))
			ub, err := ToBytes(u)
			Expect(err).To(BeNil())
			Expect(len(ub)).To(BeEquivalentTo(UUIDBytesSize))
		})
	})

})
