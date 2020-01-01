package dailybonus

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Daily Bonus Config tests", func() {

	Describe("Creating a Config", func() {
		It("should succeed", func() {
			By("Creating Config")
			config := &Config{}
			Expect(config).NotTo(BeNil())
		})
	})

})
