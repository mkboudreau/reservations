package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/mkboudreau/loggo")




func TestMonitorJsonConfig(t *testing.T) {
	testLogger := loggo.DefaultLevelLogger()
	testLogger.SetInfo()

	Convey("When doing basic unmarshalling", t, func() {

			Convey("When the json is valid, unmarshalling the byte array should be successful", func() {
					jsondata := []byte(monitorJsonTestString)

					monitors, ok := unmarshalMonitorJson(jsondata)

					testLogger.Trace(monitors)

					So(ok, ShouldBeTrue)
					So(monitors, ShouldNotBeNil)
					So(monitors.Monitors, ShouldNotBeNil)
					So(len(monitors.Monitors), ShouldBeGreaterThan, 0)
				})
			Convey("When the json is valid, unmarshalling the string should be successful", func() {
					monitors, ok := NewMonitorsFromJson(monitorJsonTestString)

					testLogger.Trace(monitors)

					So(ok, ShouldBeTrue)
					So(monitors, ShouldNotBeNil)
					So(monitors.Monitors, ShouldNotBeNil)
					So(len(monitors.Monitors), ShouldBeGreaterThan, 0)
				})

		})
	Convey("When marshalling from the stored asset file", t, func() {

			Convey("When the json is valid, unmarshalling should be successful", func() {
					monitors, ok := NewDefaultMonitors()

					testLogger.Trace(monitors)

					So(ok, ShouldBeTrue)
					So(monitors, ShouldNotBeNil)
					So(monitors.Monitors, ShouldNotBeNil)
					So(len(monitors.Monitors), ShouldBeGreaterThan, 0)
				})

		})
	Convey("When marshalling from an external json file", t, func() {

			Convey("When the json is valid, unmarshalling should be successful", func() {
					monitors, ok := NewMonitorsFromFile("../assets/monitors.json")

					testLogger.Trace(monitors)

					So(ok, ShouldBeTrue)
					So(monitors, ShouldNotBeNil)
					So(monitors.Monitors, ShouldNotBeNil)
					So(len(monitors.Monitors), ShouldBeGreaterThan, 0)
				})

		})
}



const monitorJsonTestString string = `
{
    "monitors": {
        "TestYosemite2014" : {
            "name": "TestYosemite2014",
            "emails": [ "7149146095@txt.att.net"],
            "location": 70925,
            "from": "2014-06-19",
            "to": "2014-06-24",
            "sites": [237, 235, 233, 231, 229, 227, 225, 223, 221, 219,
                217, 215, 212, 209, 205, 203, 201, 200, 198, 196, 194, 192,
                190, 188, 197 ],
            "hpsites": [197, 196],
            "hpdates": [ "2014-06-20","2014-06-23", "2014-06-24"]
        }
    }
}

`
