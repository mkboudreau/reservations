package site_test

import (
	. "github.com/mkboudreau/reservations/site"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const siteRetrRecGovTestUrl = "http://www.recreation.gov/campsiteCalendar.do?sitepage=true&parkId=70925&page=calendar&startIdx=200&calarvdate=06%2F19%2F2014&contractCode=NRSO"

var _ = Describe("Site Retrieval Tests", func() {
		Context("Retrieve HTML Test", func() {
				It("Get the start command", func() {
						html, err := RetrieveHtmlFromURL(siteRetrRecGovTestUrl)
						Expect(html).ToNot(BeNil())
						Expect(err).To(BeNil())
						Expect(html).To(ContainSubstring("</html>"))
					})
			})

	})


