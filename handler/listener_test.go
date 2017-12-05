package handler_test

import (
	"bytes"

	"github.com/jamesjoshuahill/language/handler"
	"github.com/jamesjoshuahill/language/handler/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	It("has a language handler", func() {
		recorder := new(fakes.FakeRecorder)
		listener := handler.New(recorder)
		reader := bytes.NewReader([]byte("hello"))
		conn := new(fakes.FakeConn)
		conn.ReadStub = reader.Read

		listener.Handle(conn)

		Expect(conn.ReadCallCount()).To(Equal(2))
		Expect(recorder.RecordCallCount()).To(Equal(1))
		Expect(recorder.RecordArgsForCall(0)).To(Equal("hello"))
	})
})
