package stats_test

import (
	"github.com/jamesjoshuahill/language/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stats", func() {
	It("can count words", func() {
		stats := stats.New()

		stats.Record("here are some more words")

		Expect(stats.Summary().Count).To(Equal(5))
	})

	It("counts each word once", func() {
		stats := stats.New()

		stats.Record("here are some more words")
		stats.Record("and here are some other words")

		Expect(stats.Summary().Count).To(Equal(7))
	})
})
