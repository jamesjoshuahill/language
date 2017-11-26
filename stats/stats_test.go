package stats_test

import (
	"github.com/jamesjoshuahill/language/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stats", func() {
	It("can count words", func() {
		stats := stats.NewStats()

		stats.Record("here are some more words")

		Expect(stats.Summary().Count).To(Equal(5))
	})

	It("counts each word once", func() {
		stats := stats.NewStats()

		stats.Record("here are some more words")
		stats.Record("and here are some other words")

		Expect(stats.Summary().Count).To(Equal(7))
	})

	It("knows the top five words", func() {
		stats := stats.NewStats()

		stats.Record("here are some more words")

		Expect(stats.Summary().Top5Words).
			To(ConsistOf("here", "are", "some", "more", "words"))
	})

	It("updates the list of top five words", func() {
		stats := stats.NewStats()

		stats.Record("here are some more words")
		stats.Record("six six six six six six")
		stats.Record("five five five five five")
		stats.Record("four four four four")
		stats.Record("three three three")
		stats.Record("two two")

		Expect(stats.Summary().Top5Words).
			To(ConsistOf("six", "five", "four", "three", "two"))
	})

	It("knows the top five letters", func() {
		stats := stats.NewStats()

		stats.Record("a bb ccc dddd eeeee ffffff")

		Expect(stats.Summary().Top5Letters).
			To(ConsistOf("b", "c", "d", "e", "f"))
	})
})
