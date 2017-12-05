package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"

	"github.com/jamesjoshuahill/language/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Server", func() {
	const (
		listenerPort = 5555
		apiPort      = 8080
	)

	AfterEach(func() {
		gexec.KillAndWait()
	})

	It("listens on ports 5555 and 8080 by default", func() {
		session := startServer(listenerPort, apiPort)

		Expect(session.Out.Contents()).To(ContainSubstring("Listening on port 5555..."))
		Expect(session.Out.Contents()).To(ContainSubstring("Serving HTTP on port 8080..."))
		Consistently(session).ShouldNot(gexec.Exit())
	})

	It("listens for natural language", func() {
		session := startServer(listenerPort)

		sendLanguage("here are some more words", listenerPort)

		Eventually(session.Out).Should(gbytes.Say("received 'here'"))
	})

	It("can listen for natural language on a custom port", func() {
		cmd := exec.Command(serverBinaryPath, "-port", "1234")

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(listeningOn(1234)).Should(BeTrue())
		Expect(session.Out).To(gbytes.Say("Listening on port 1234..."))
	})

	It("responds to GET /stats", func() {
		startServer(apiPort)

		response, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/stats", apiPort))

		Expect(err).NotTo(HaveOccurred())
		defer response.Body.Close()
		stats, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(stats).To(MatchJSON(`{
			"count": 0,
			"top5words": [],
			"top5letters": []
		}`))
	})

	It("can serve HTTP requests on a custom port", func() {
		cmd := exec.Command(serverBinaryPath, "-apiPort", "1234")

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(listeningOn(1234)).Should(BeTrue())
		Expect(session.Out).To(gbytes.Say("Serving HTTP on port 1234..."))
		_, err = http.DefaultClient.Get("http://localhost:1234/stats")
		Expect(err).NotTo(HaveOccurred())
	})

	It("responds to GET /stats with language stats", func() {
		startServer(listenerPort, apiPort)
		sendLanguage("a a at at hat hat chat chat match match", listenerPort)

		response, err := http.DefaultClient.Get("http://localhost:8080/stats")

		Expect(err).NotTo(HaveOccurred())
		defer response.Body.Close()
		var actual stats.Summary
		Expect(json.NewDecoder(response.Body).Decode(&actual)).To(Succeed())
		Expect(actual.Count).To(Equal(5))
		Expect(actual.Top5Words).To(ConsistOf("a", "at", "hat", "chat", "match"))
		Expect(actual.Top5Letters).To(ConsistOf("a", "t", "h", "c", "m"))
	})
})

func startServer(ports ...int) *gexec.Session {
	cmd := exec.Command(serverBinaryPath)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	for _, port := range ports {
		Eventually(listeningOn(port)).Should(BeTrue())
	}
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

func sendLanguage(language string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	Expect(err).NotTo(HaveOccurred())

	_, err = conn.Write([]byte(language))

	Expect(err).NotTo(HaveOccurred())
	conn.Close()
}
