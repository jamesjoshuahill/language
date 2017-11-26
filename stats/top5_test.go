package stats_test

import (
	"github.com/jamesjoshuahill/language/stats"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Top5", func() {
	It("starts with none", func() {
		top5 := stats.NewTop5()

		Expect(top5.List()).To(BeEmpty())
	})

	It("discards zero scores", func() {
		top5 := stats.NewTop5()

		top5.Submit("zero", 0)

		Expect(top5.List()).To(BeEmpty())
	})

	It("records five entries", func() {
		top5 := stats.NewTop5()

		top5.Submit("one", 1).
			Submit("two", 2).
			Submit("three", 3).
			Submit("four", 4).
			Submit("five", 5)

		Expect(top5.List()).To(ConsistOf("one", "two", "three", "four", "five"))
	})

	It("knows the lowest entry", func() {
		top5 := stats.NewTop5()

		top5.Submit("one", 1).
			Submit("two", 2).
			Submit("three", 3).
			Submit("four", 4).
			Submit("five", 5)

		entry, score := top5.Lowest()
		Expect(entry).To(Equal("one"))
		Expect(score).To(Equal(1))
	})

	It("discards the lowest entry when there are more than five entries", func() {
		top5 := stats.NewTop5()

		top5.Submit("one", 1).
			Submit("seven", 7).
			Submit("two", 2).
			Submit("three", 3).
			Submit("five", 5).
			Submit("four", 4).
			Submit("six", 6)

		Expect(top5.List()).To(ConsistOf("three", "four", "five", "six", "seven"))
	})
})
