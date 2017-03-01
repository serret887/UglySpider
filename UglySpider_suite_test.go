package UglySpider_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUglySpider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UglySpider Suite")
}
