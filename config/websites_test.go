package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/mkboudreau/loggo")




func TestWebsiteJsonConfig(t *testing.T) {
	testLogger := loggo.DefaultLevelLogger()
	testLogger.SetDebug()

	Convey("When doing basic unmarshalling", t, func() {

			Convey("When the json is valid, unmarshalling the byte array should be successful", func() {
					jsondata := []byte(websiteJsonTestString)

					websites, ok := unmarshalWebsiteJson(jsondata)

					testLogger.Debug(websites)

					So(ok, ShouldBeTrue)
					So(websites, ShouldNotBeNil)
					So(websites.Websites, ShouldNotBeNil)
					So(len(websites.Websites), ShouldBeGreaterThan, 0)
				})
			Convey("When the json is valid, unmarshalling the string should be successful", func() {
					websites, ok := NewWebsitesFromJson(websiteJsonTestString)

					testLogger.Debug(websites)

					So(ok, ShouldBeTrue)
					So(websites, ShouldNotBeNil)
					So(websites.Websites, ShouldNotBeNil)
					So(len(websites.Websites), ShouldBeGreaterThan, 0)
				})

		})
	Convey("When marshalling from the stored asset file", t, func() {

			Convey("When the json is valid, unmarshalling should be successful", func() {
					websites, ok := NewDefaultWebsites()

					testLogger.Debug(websites)

					So(ok, ShouldBeTrue)
					So(websites, ShouldNotBeNil)
					So(websites.Websites, ShouldNotBeNil)
					So(len(websites.Websites), ShouldBeGreaterThan, 0)
				})

		})
	Convey("When marshalling from an external json file", t, func() {

			Convey("When the json is valid, unmarshalling should be successful", func() {
					websites, ok := NewWebsitesFromFile("../assets/websites.json")

					testLogger.Debug(websites)

					So(ok, ShouldBeTrue)
					So(websites, ShouldNotBeNil)
					So(websites.Websites, ShouldNotBeNil)
					So(len(websites.Websites), ShouldBeGreaterThan, 0)
				})

		})
}



const websiteJsonTestString string = `
{
    "websites": [
        {
            "url": "http://www.recreation.gov/campsiteCalendar.do?",
            "static_params" : {
                "contract_code" : "NRSO",
                "page" : "calendar",
                "sitepage":"true"
            },
            "dynamic_params" : {
                "site" : "parkId",
                "start_index" : "startIdx",
                "start_date":"calarvdate"
            }
        }
    ]
}


`
