package userdata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/ginkgo/reporters"
)

func TestGuest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "User Data Suite",
		[]Reporter{reporters.NewJUnitReporter("testUserData.xml")})
}
