package site

import (
	"github.com/mkboudreau/loggo"
 	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"time"
)

var logger *loggo.LevelLogger = loggo.DefaultLevelLogger()

type SiteParser struct {
	ExpectedStartDay time.Time
}

func NewSiteParser(expectedStartDay time.Time) *SiteParser {
	return &SiteParser{ExpectedStartDay:expectedStartDay}
}

func (parser *SiteParser) ParseHtmlFromUrl(url string) ([]Site, error) {
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument(url); e != nil {
		logger.Fatal(e)
	}


	return parser.parseDocument(doc)
}
func (parser *SiteParser) ParseHtmlString(html string) ([]Site, error) {
	var doc *goquery.Document
	var e error

	reader := strings.NewReader(html)
	if doc, e = goquery.NewDocumentFromReader(reader); e != nil {
		logger.Fatal(e)
	}

	return parser.parseDocument(doc)
}

func (parser *SiteParser) parseDocument(doc *goquery.Document) (sites []Site, err error) {
	sites = make([]Site, 0)

	doc.Find("#calendar").Find("tr").Each(func( trIndex int, tr *goquery.Selection) {
		logger.Trace("Found Calendar Row")
		tr.Find(".calendar").Each(func(headerIndex int, th *goquery.Selection) {
			logger.Trace("Found Calendar Headers")
			// do nothing right now
		})

		tr.Find(".sn").Each(func(siteIndex int, sn *goquery.Selection) {
			logger.Trace("Found site section")
			siteNumberDivSelection := sn.First()
			siteNumberDivSelection.Find(".siteListLabel").Each(func(siteLabelIndex int, sitelist *goquery.Selection) {
				siteNumber := sitelist.First().Text()
				logger.Trace("Found site number [",siteNumber,"]")

				// CREATE THE SITE OBJECT
				//final Site site = new Site(Integer.parseInt(siteNumber));
				site := Site{}
				site.SiteNumber, _ = strconv.Atoi(siteNumber)

				tr.Find(".status").Each(func(siteDayOffset int, status *goquery.Selection) {
					statusNode := status.Get(0)
					hoursInOneDay := 24
					offsetInHours := hoursInOneDay * siteDayOffset
					durationString := strconv.Itoa(offsetInHours) + "h"
					dayOffsetDuration, _ :=  time.ParseDuration(durationString)

					for _, statusAttribute := range statusNode.Attr {
						logger.Trace("Found attributes [", statusAttribute, "]")
						if (statusAttribute.Key == "class") {
							classes := strings.Split(statusAttribute.Val, " ")
							logger.Trace("  -- classes: [",classes,"]")
							for _, statusCode := range classes {
								isValid := IsValidSiteAvailabilityCode(statusCode)
								logger.Trace("  -- availability: [",statusCode,"]; Valid Code?",isValid)

								if isValid {
									logger.Trace("  -- processing day --")
									statusAvailabilityCode := GetSiteAvailabilityCodeForLetter(statusCode)
									siteDay := SiteDay{}
									siteDay.SiteAvailability = statusAvailabilityCode
									logger.Trace("    - adding ", dayOffsetDuration, " to expected start day ", parser.ExpectedStartDay)
									siteDay.Day = parser.ExpectedStartDay.Add(dayOffsetDuration)
									logger.Trace("    - resulting day ", siteDay.Day)


									// build a DAY object
									// set the OFFSET
									// set the AVAILABILITY STATUS CODE
									// ADD THE DAY TO THE SITE

									site.SiteDays = append(site.SiteDays, siteDay)
								}
							}
						}


					}

				});
				// ADD SITE TO THE SITE ARRAY/LIST
				//sites.add(site);
				sites=append(sites, site)
			})
		})
	})

	return sites, err
}
