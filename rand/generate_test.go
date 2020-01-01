package rand

import (
	"encoding/base64"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rand generate tests", func() {

	Describe("GenerateRandomBytes with size 1", func() {
		It("should generate a byte array of size 1", func() {
			By("Generating random bytes")
			bytes, err := GenerateRandomBytes(1)
			Expect(err).To(BeNil())
			Expect(bytes).NotTo(BeNil())
			Expect(len(bytes)).To(BeEquivalentTo(1))
		})
	})

	Describe("GenerateRandomString with size 1", func() {
		It("should generate a string of size 1", func() {
			By("Generating random string")
			str, err := GenerateRandomString(1)
			Expect(err).To(BeNil())
			expectedLen := len(base64.URLEncoding.EncodeToString([]byte{0}))
			Expect(len(str)).To(BeEquivalentTo(expectedLen))
		})
	})

})
