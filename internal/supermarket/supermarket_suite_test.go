package supermarket_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSupermarket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Supermarket Suite")
}
