package dailybonus_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestGuest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Daily Bonus Suite",
		[]Reporter{reporters.NewJUnitReporter("testDailyBonus.xml")})
}
