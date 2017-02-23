package pitcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPitcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pitcher Suite")
}
