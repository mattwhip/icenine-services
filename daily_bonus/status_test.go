package dailybonus

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Daily Bonus Status tests", func() {

	Describe("Creating a Status", func() {
		It("should succeed", func() {
			By("Creating Status")
			status := &Status{}
			Expect(status).NotTo(BeNil())
		})
	})

})
