package rand_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/ginkgo/reporters"
)

func TestGuest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Rand Suite",
		[]Reporter{reporters.NewJUnitReporter("testRand.xml")})
}
