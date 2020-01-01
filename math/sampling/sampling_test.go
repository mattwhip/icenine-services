package sampling

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sampling tests", func() {

	Describe("Creating a Population with one item", func() {
		It("should succeed", func() {
			By("Creating Population")
			population, err := NewPopulation([]interface{}{"a"}, []int64{1})
			Expect(err).NotTo(HaveOccurred())
			Expect(population).NotTo(BeNil())
		})
	})

	Describe("Sampling a Population with one item", func() {
		It("should return the one item", func() {
			By("Creating Population")
			population, err := NewPopulation([]interface{}{"a"}, []int64{1})
			Expect(err).NotTo(HaveOccurred())
			Expect(population).NotTo(BeNil())

			item, err := population.Sample()
			Expect(err).NotTo(HaveOccurred())
			Expect(item.Value).To(BeEquivalentTo("a"))
		})
	})

})
