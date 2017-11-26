package main_test

import (
	"fmt"
	"net"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Server", func() {
	var session *gexec.Session

	BeforeEach(func() {
		session = startServer(5555)
	})

	AfterEach(func() {
		gexec.KillAndWait()
	})

	It("listens on port 5555", func() {
		Expect(session.Out).To(gbytes.Say("Starting server on port 5555..."))
		Consistently(session).ShouldNot(gexec.Exit())
	})
})

func startServer(port int) *gexec.Session {
	cmd := exec.Command(serverBinaryPath)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	Eventually(listeningOn(port)).Should(BeTrue())
	return session
}

func listeningOn(port int) func() bool {
	return func() bool {
		conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return false
		}

		conn.Close()
		return true
	}
}
