package lightswitchcontroller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLightswitchController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lightswitch Controller")
}
