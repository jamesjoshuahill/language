package main_test

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
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

		_, err = conn.Write([]byte("here are some more words"))
		conn.Close()

		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("received 'here are some more words'"))
	})

	It("responds to GET /stats", func() {
		startServer(port)

		response, err := http.DefaultClient.Get("http://localhost:8080/stats")

		Expect(err).NotTo(HaveOccurred())
		defer response.Body.Close()
		stats, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(stats).To(MatchJSON(`{
			"count": 5,
			"top5words": ["here", "are", "some", "more", "words"],
			"top5letters": ["e", "r", "o", "s", "h"]
		}`))
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
