package location

import (
	"errors"
	"github.com/mkboudreau/loggo"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestLocationAPI(t *testing.T) {
	testLogger := loggo.DefaultLevelLogger()
	testLogger.SetInfo()

	Convey("When a valid xml string is parsed", t, func() {
		Convey("It should respond with Locations", func() {
			xml := []byte(locationListXml)
			locations, ok := unmarshallXmlResponseIntoLocations(xml)

			So(ok, ShouldBeTrue)
			So(locations, ShouldNotBeNil)
		})
	})

	Convey("When an invalid xml string is parsed", t, func() {
		Convey("It should return false", func() {
			xml := []byte(veryInvalidXml)
			locations, ok := unmarshallXmlResponseIntoLocations(xml)

			So(ok, ShouldBeFalse)
			So(locations, ShouldBeNil)
		})
	})

	Convey("When a user asks for locations from an invalid server", t, func() {
		mockedLocationService := newRecreationGovService()
		mockedLocationService.serviceUrl = "http://localhost:9999/hello"

		Convey("It should return a result of false", func() {
			locations, ok := mockedLocationService.GetAllLocations()

			So(ok, ShouldBeFalse)
			So(locations, ShouldBeNil)
		})

	})

	Convey("When a user asks for locations from the wrong server", t, func() {
		mockedLocationService := newRecreationGovService()
		mockedLocationService.serviceUrl = "http://yahoo.com"

		Convey("It should return a result of false", func() {
			locations, ok := mockedLocationService.GetAllLocations()

			So(ok, ShouldBeFalse)
			So(locations, ShouldBeNil)
		})

	})

	Convey("When there is a problem reading the response from the server", t, func() {
		mockedResponse := buildMockedHttpResponse()

		//func getLocationsFromHttpResponse(resp http.Response) (*Locations, bool) {

		Convey("It should return a result of false", func() {
			locations, ok := getLocationsFromHttpResponse(mockedResponse)

			So(ok, ShouldBeFalse)
			So(locations, ShouldBeNil)
		})

	})

	Convey("When a user asks for locations", t, func() {
		mockedLocationService := newRecreationGovService()
		mockedLocationService.cachedData = buildMockLocations()

		Convey("It should respond with all locations", func() {
			locations, ok := mockedLocationService.GetAllLocations()

			So(ok, ShouldBeTrue)
			So(locations, ShouldNotBeNil)
			So(len(locations.Locations), ShouldEqual, 3)
		})

	})

	Convey("When a user asks for a single location by Id", t, func() {
		mockedLocationService := newRecreationGovService()
		mockedLocationService.cachedData = buildMockLocations()

		Convey("When the ID is known, it should respond with a single location object and a boolean true", func() {
			knownId := "70925"
			location, ok := mockedLocationService.GetLocationById(knownId)

			So(ok, ShouldBeTrue)
			So(location, ShouldNotBeNil)
			So(location.Name, ShouldEqual, "Test 70925")
			So(location.LegacyID, ShouldEqual, knownId)

		})
		Convey("When the ID is not known, it should respond with a boolean false", func() {
			unknownId := "222222222222222222"
			location, ok := mockedLocationService.GetLocationById(unknownId)

			So(ok, ShouldBeFalse)
			So(location, ShouldBeNil)
		})

		Convey("When it is the first time this method is called and it cannot get the locations, it should respond with a boolean false", func() {

			mockedProblemLocationService := newRecreationGovService()
			mockedProblemLocationService.serviceUrl = "http://yahoo.com"

			unknownId := "70925"
			location, ok := mockedProblemLocationService.GetLocationById(unknownId)

			So(ok, ShouldBeFalse)
			So(location, ShouldBeNil)
		})

	})

	Convey("Full Integration Tests", t, func() {
		locationService := NewLocationService()

		Convey("It should respond with all locations", func() {
			locations, ok := locationService.GetAllLocations()

			So(ok, ShouldBeTrue)
			So(locations, ShouldNotBeNil)
			So(len(locations.Locations), ShouldBeGreaterThan, 10)
		})

		Convey("When the ID is known, it should respond with a single location object and a boolean true", func() {
			knownId := "70925"
			location, ok := locationService.GetLocationById(knownId)

			So(ok, ShouldBeTrue)
			So(location, ShouldNotBeNil)
			So(location.Name, ShouldStartWith, "UPPER PINES")
			So(location.LegacyID, ShouldEqual, knownId)

		})
		Convey("When the ID is not known, it should respond with a boolean false", func() {
			unknownId := "222222222222222222"
			location, ok := locationService.GetLocationById(unknownId)

			So(ok, ShouldBeFalse)
			So(location, ShouldBeNil)
		})

	})

}

func buildMockLocations() *Locations {
	return &Locations{
		Locations: []*Location{
			buildMockLocation("70925"),
			buildMockLocation("88988"),
			buildMockLocation("88888"),
		},
	}
}

func buildMockLocation(id string) *Location {
	return &Location{
		ID:            id,
		Name:          "Test " + id,
		LegacyID:      id,
		OrgFacilityID: "AN" + id,
	}
}

type mockreadcloser struct{}

func (reader *mockreadcloser) Read(p []byte) (int, error) {
	return 0, errors.New("Mock Read Error")
}
func (closer *mockreadcloser) Close() error {
	return errors.New("Mock Close Error")
}

func buildMockedHttpResponse() *http.Response {
	resp := &http.Response{
		Body: new(mockreadcloser),
	}

	return resp
}

const locationListXml = `
<arc:RecElements xmlns:arc="http://www.recreation.gov/architecture/">
<arc:Facility xmlns:arc="http://www.recreation.gov/architecture/">
	<arc:FacilityID>232447</arc:FacilityID>
	<arc:OrgFacilityID>AN370925</arc:OrgFacilityID>
	<arc:FacilityName>UPPER PINES</arc:FacilityName>
	<arc:FacilityTypeDescription>Camping</arc:FacilityTypeDescription>
	<arc:FacilityPhone/>
	<arc:FacilityDescription>
		<h2>Overview</h2>Upper Pines Campground is located in breathtaking Yosemite National Park in Central Californias rugged Sierra Nevada Mountain Range at an elevation of 4,000 feet. The site is situated in the heart of Yosemite Valley, an awe-inspiring landscape containing many of the famous features for which Yosemite National Park is known.<br /><br />Within Yosemite, visitors can gaze upon waterfalls, sheer granite cliffs, deep valleys, grand meadows, ancient giant sequoias, vast wilderness areas and much more. <h4>Natural Features:</h4>Yosemite Valley is forested with a diverse mix of California black oak, Ponderosa pine, incense-cedar, White fir, and Whiteleaf Manzanita, which offers a pleasant combination of sun and shade. Through the trees, Yosemites telltale granite cliffs peek through, and the gentle, refreshing Merced River flows nearby.<h4>Recreation:</h4>Yosemites trails, cliffs, roads and rivers provide endless recreational activities for any kind of visitor. Hiking, biking, rock climbing, fishing, horseback riding, rafting, auto touring, cross-country skiing and photography are all very popular activities within the park. <br /><br />Hiking trails range from the easy paved, two-mile roundtrip walk to Mirror Lake/Meadow, to the grueling but rewarding 14- to 16-mile trek to Half Dome (permit required). Both of these trails begin near the Pines Campgrounds, and there are many additional trailheads within Yosemite Valley. <br /><br />Rafting the Merced River is a fun way to cool down on a summer day when water levels are sufficient. Yosemite Valley also offers numerous guided bus tours, educational programs, museums, ranger-led activities and an art center with workshops.<h4>Facilities:</h4>Upper Pines Campground is the largest of the three reservation campgrounds in Yosemite Valley It offers paved roads and parking spurs, flush toilets, and drinking water. Each campsite contains a picnic table, fire ring, and a food storage locker. Yosemites free shuttle bus stops near the campground entrance.<br /><br />Nearby historic Curry Village offers a general store, restaurants and bar, amphitheater, coin showers, and tours and activities desk.<h4>Nearby Attractions:</h4>Glacier Point offers sweeping views of Yosemite Valley, Little Yosemite Valley, Half Dome, Vernal and Nevada Falls, and Clouds Rest, among other notable landmarks. <br /><br />The Mariposa Grove of Giant Sequoias is a must-see. A guided bus tour (tickets must be purchased) and a network of hiking trails leads to stately trees with names like Grizzly Giant, Faithful Couple, and California Tunnel Tree. <br /><br />The Tioga Road offers a 39-mile scenic drive past forests, meadows, lakes and granite domes. Beautiful Hetch Hetchy Reservoir is home to spectacular scenery and numerous wilderness trailheads.
	</arc:FacilityDescription>
	<arc:FacilityDirections>
		Take Highway 41 north from Fresno, Highway 140 east from Merced, or Highway 120 east from Manteca into Yosemite National Park. Follow signs to Yosemite Valley and the campgrounds.
	</arc:FacilityDirections>
	<arc:FacilityEmail/>
	<arc:FacilityLatitude>37.736111111111</arc:FacilityLatitude>
	<arc:FacilityLongitude>-119.562500000000</arc:FacilityLongitude>
	<arc:FacilityAdaAccess/>
	<arc:FacilityUseFeeDescription/>
	<arc:Enabled>1</arc:Enabled>
	<arc:LastUpdatedBy>1101</arc:LastUpdatedBy>
	<arc:LegacyFacilityID>70925</arc:LegacyFacilityID>
</arc:Facility>
<arc:Facility xmlns:arc="http://www.recreation.gov/architecture/">
	<arc:FacilityID>77777</arc:FacilityID>
	<arc:OrgFacilityID>AN88888</arc:OrgFacilityID>
	<arc:FacilityName>SMOKEY REDWOODS</arc:FacilityName>
	<arc:FacilityTypeDescription>Camping</arc:FacilityTypeDescription>
	<arc:FacilityPhone/>
	<arc:FacilityDescription>
		<h2>Overview</h2>Upper Pines Campground is located in breathtaking Yosemite National Park in Central Californias rugged Sierra Nevada Mountain Range at an elevation of 4,000 feet. The site is situated in the heart of Yosemite Valley, an awe-inspiring landscape containing many of the famous features for which Yosemite National Park is known.<br /><br />Within Yosemite, visitors can gaze upon waterfalls, sheer granite cliffs, deep valleys, grand meadows, ancient giant sequoias, vast wilderness areas and much more. <h4>Natural Features:</h4>Yosemite Valley is forested with a diverse mix of California black oak, Ponderosa pine, incense-cedar, White fir, and Whiteleaf Manzanita, which offers a pleasant combination of sun and shade. Through the trees, Yosemites telltale granite cliffs peek through, and the gentle, refreshing Merced River flows nearby.<h4>Recreation:</h4>Yosemites trails, cliffs, roads and rivers provide endless recreational activities for any kind of visitor. Hiking, biking, rock climbing, fishing, horseback riding, rafting, auto touring, cross-country skiing and photography are all very popular activities within the park. <br /><br />Hiking trails range from the easy paved, two-mile roundtrip walk to Mirror Lake/Meadow, to the grueling but rewarding 14- to 16-mile trek to Half Dome (permit required). Both of these trails begin near the Pines Campgrounds, and there are many additional trailheads within Yosemite Valley. <br /><br />Rafting the Merced River is a fun way to cool down on a summer day when water levels are sufficient. Yosemite Valley also offers numerous guided bus tours, educational programs, museums, ranger-led activities and an art center with workshops.<h4>Facilities:</h4>Upper Pines Campground is the largest of the three reservation campgrounds in Yosemite Valley It offers paved roads and parking spurs, flush toilets, and drinking water. Each campsite contains a picnic table, fire ring, and a food storage locker. Yosemites free shuttle bus stops near the campground entrance.<br /><br />Nearby historic Curry Village offers a general store, restaurants and bar, amphitheater, coin showers, and tours and activities desk.<h4>Nearby Attractions:</h4>Glacier Point offers sweeping views of Yosemite Valley, Little Yosemite Valley, Half Dome, Vernal and Nevada Falls, and Clouds Rest, among other notable landmarks. <br /><br />The Mariposa Grove of Giant Sequoias is a must-see. A guided bus tour (tickets must be purchased) and a network of hiking trails leads to stately trees with names like Grizzly Giant, Faithful Couple, and California Tunnel Tree. <br /><br />The Tioga Road offers a 39-mile scenic drive past forests, meadows, lakes and granite domes. Beautiful Hetch Hetchy Reservoir is home to spectacular scenery and numerous wilderness trailheads.
	</arc:FacilityDescription>
	<arc:FacilityDirections>
		Take Highway 41 north from Fresno, Highway 140 east from Merced, or Highway 120 east from Manteca into Yosemite National Park. Follow signs to Yosemite Valley and the campgrounds.
	</arc:FacilityDirections>
	<arc:FacilityEmail/>
	<arc:FacilityLatitude>37.736111111111</arc:FacilityLatitude>
	<arc:FacilityLongitude>-119.562500000000</arc:FacilityLongitude>
	<arc:FacilityAdaAccess/>
	<arc:FacilityUseFeeDescription/>
	<arc:Enabled>1</arc:Enabled>
	<arc:LastUpdatedBy>1101</arc:LastUpdatedBy>
	<arc:LegacyFacilityID>88888</arc:LegacyFacilityID>
</arc:Facility>
</arc:RecElements>
`

const veryInvalidXml = `
<<<DdfdfkdkjflJL!!!.>>>
`
