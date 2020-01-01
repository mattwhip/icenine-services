package middleware

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JWT verification tests", func() {

	Describe("Creating a JWTVerification middleware", func() {
		It("should succeed", func() {
			By("Creating middlware")
			middlware := JWTVerification("dummy-signing-secret")
			Expect(middlware).NotTo(BeNil())
		})
	})

})
