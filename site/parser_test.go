package site_test

import (
	. "github.com/mkboudreau/reservations/site"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/mkboudreau/loggo"
	"time"
)

const siteParserRecGovTestUrl = "http://www.recreation.gov/campsiteCalendar.do?sitepage=true&parkId=70925&page=calendar&startIdx=200&calarvdate=06%2F19%2F2014&contractCode=NRSO"


var _ = Describe("Site Parsing Tests", func() {
		var logger *loggo.LevelLogger = loggo.DefaultLevelLogger()
		logger.SetTrace()
		Context("Parse and Print HTML", func() {
				t := time.Date(2014, time.June, 19, 12, 0, 0, 0, time.UTC)
				parser := NewSiteParser(t)
				siteSlice, err := parser.ParseHtmlFromUrl(siteParserRecGovTestUrl)
				logger.Trace(siteSlice)
				Expect(siteSlice).ToNot(BeNil())
				Expect(err).To(BeNil())
				/*
				It("Get the start command", func() {
						html, err := remoteSite.RetrieveHtml()
						Expect(html).ToNot(BeNil())
						Expect(err).To(BeNil())
						Expect(html).To(ContainSubstring("</html>"))
					})
					*/
			})

	})


const HtmlTest = `

<html>
<head><title>Hello</title>

</head>

<body>

<div id="id1" class="someclass"> <h1>Test H1</h1></div>

<div></div>

<div>       </div>
</body>
</html>


`

