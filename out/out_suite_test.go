package main_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var outPath string

var _ = BeforeSuite(func() {
	var err error

	if _, err = os.Stat("/opt/resource/out"); err == nil {
		outPath = "/opt/resource/out"
	} else {
		outPath, err = gexec.Build("github.com/concourse/time-resource/out")
		Expect(err).NotTo(HaveOccurred())
	}
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestOut(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Out Suite")
}