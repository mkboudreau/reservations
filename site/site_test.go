package site_test

import (
	. "github.com/mkboudreau/reservations/site"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reservations Site Test Suite")
}

var _ = Describe("Test Site Availability for Code Letter", func() {
		It("Get a", func() {
				code := GetSiteAvailabilityCodeForLetter("a")
				Expect(code).To(Equal(SiteAvailable))
			})
		It("Get A", func() {
				code := GetSiteAvailabilityCodeForLetter("A")
				Expect(code).To(Equal(SiteAvailable))
			})

	})
