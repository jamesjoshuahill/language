package handler_test

import (
	"bytes"
	"io/ioutil"

	"github.com/jamesjoshuahill/language/handler"
	"github.com/jamesjoshuahill/language/handler/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	It("calls the recorder with the word received", func() {
		recorder := new(fakes.FakeRecorder)
		listener := handler.New(recorder)
		conn := ioutil.NopCloser(bytes.NewReader([]byte("hello")))

		listener.Handle(conn)

		Expect(recorder.RecordCallCount()).To(Equal(1))
		Expect(recorder.RecordArgsForCall(0)).To(Equal("hello"))
	})

	It("calls the recorder with each word received", func() {
		recorder := new(fakes.FakeRecorder)
		listener := handler.New(recorder)
		conn := ioutil.NopCloser(bytes.NewReader([]byte("hello world")))

		listener.Handle(conn)

		Expect(recorder.RecordCallCount()).To(Equal(2))
		Expect(recorder.RecordArgsForCall(0)).To(Equal("hello"))
		Expect(recorder.RecordArgsForCall(1)).To(Equal("world"))
	})
})
