package middleware

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiPopTransaction tests", func() {

	Describe("Creating a MultiPopTransaction middleware", func() {
		It("should succeed", func() {
			By("Creating middlware")
			middlware := MultiPopTransaction(nil, "")
			Expect(middlware).NotTo(BeNil())
		})
	})

})
