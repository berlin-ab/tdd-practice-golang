package rockpaperscissors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRockPaperScissorsGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RockPaperScissorsGolang Suite")
}
