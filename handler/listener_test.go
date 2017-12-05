package handler_test

import (
	"io"

	"github.com/jamesjoshuahill/language/handler"
	"github.com/jamesjoshuahill/language/handler/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	It("has a language handler", func() {
		recorder := new(fakes.FakeRecorder)
		listener := handler.New(recorder)
		conn := new(fakes.FakeConn)
		conn.ReadReturns(0, io.EOF)

		listener.Handle(conn)

		Expect(conn.ReadCallCount()).To(Equal(1))
	})
})
