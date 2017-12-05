package stats_test

import (
	"github.com/jamesjoshuahill/language/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stats", func() {
	It("starts with an empty summary", func() {
		stats := stats.NewStats()

		Expect(stats.Summary().Count).To(BeZero())
		Expect(stats.Summary().Top5Words).To(BeEmpty())
		Expect(stats.Summary().Top5Letters).To(BeEmpty())
	})

	It("can record words", func() {
		stats := stats.NewStats()

		stats.Record("hello")
		stats.Record("world")

		Expect(stats.Summary().Count).To(Equal(2))
	})

	It("counts each word once", func() {
		stats := stats.NewStats()

		stats.Record("hello")
		stats.Record("hello")
		stats.Record("world")

		Expect(stats.Summary().Count).To(Equal(2))
	})

	It("knows the top five words", func() {
		stats := stats.NewStats()

		stats.Record("here")
		stats.Record("are")
		stats.Record("some")
		stats.Record("more")
		stats.Record("words")

		Expect(stats.Summary().Top5Words).
			To(ConsistOf("here", "are", "some", "more", "words"))
	})

	It("updates the list of top five words", func() {
		stats := stats.NewStats()

		stats.Record("here")
		stats.Record("are")
		stats.Record("some")
		stats.Record("more")
		stats.Record("words")

		stats.Record("two")
		stats.Record("two")

		stats.Record("three")
		stats.Record("three")

		stats.Record("four")
		stats.Record("four")

		stats.Record("five")
		stats.Record("five")

		stats.Record("six")
		stats.Record("six")

		Expect(stats.Summary().Top5Words).
			To(ConsistOf("six", "five", "four", "three", "two"))
	})

	It("knows the top five letters", func() {
		stats := stats.NewStats()

		stats.Record("a")
		stats.Record("bb")
		stats.Record("ccc")
		stats.Record("dddd")
		stats.Record("eeeee")
		stats.Record("ffffff")

		Expect(stats.Summary().Top5Letters).
			To(ConsistOf("b", "c", "d", "e", "f"))
	})
})
