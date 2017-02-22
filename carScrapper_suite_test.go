package carScrapper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCarScrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CarScrapper Suite")
}
