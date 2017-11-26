package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var serverBinaryPath string

func TestLanguage(t *testing.T) {
	BeforeSuite(func() {
		var err error
		serverBinaryPath, err = gexec.Build("github.com/jamesjoshuahill/language")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}
