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
	const port = 5555

	AfterEach(func() {
		gexec.KillAndWait()
	})

	It("listens on port 5555 by default", func() {
		session := startServer(port)

		Expect(session.Out).To(gbytes.Say("Listening on port 5555..."))
		Consistently(session).ShouldNot(gexec.Exit())
	})

	It("listens on a custom port", func() {
		cmd := exec.Command(serverBinaryPath, "-port", "1234")

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(listeningOn(1234)).Should(BeTrue())
		Expect(session.Out).To(gbytes.Say("Listening on port 1234..."))
	})

	It("accepts arbitrary natural language", func() {
		session := startServer(port)
		conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
		Expect(err).NotTo(HaveOccurred())

		_, err = conn.Write([]byte("here are some words"))
		conn.Close()

		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("received 'here are some words'"))
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
