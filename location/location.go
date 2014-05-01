package location

import (
	"fmt"
	"encoding/xml"
	"github.com/mkboudreau/loggo"
	"io/ioutil"
	"net/http"
)

var locationLogger *loggo.LevelLogger = loggo.DefaultLevelLogger()
var recreationGovService *recreationGov = newRecreationGovService()

type LocationService interface {
	GetAllLocations() (*Locations, bool)
	GetLocationById(facilityId string) (*Location, bool)
}

type Locations struct {
	Locations []*Location `xml:"Facility"`
}
type Location struct {
	Name            string `xml:"FacilityName"`
	ID              string `xml:"FacilityID"`
	LegacyID        string `xml:"LegacyFacilityID"`
	OrgFacilityID   string `xml:"OrgFacilityID"`
	TypeDescription string `xml:"FacilityTypeDescription"`
	Description     string `xml:"FacilityDescription"`
	Latitude        string `xml:"FacilityLatitude"`
	Longitude       string `xml:"FacilityLongitude"`
	Enabled         string `xml:"Enabled"`
}


func NewLocationService() LocationService {
	return recreationGovService
}

func (location *Location) String() string {
	return fmt.Sprintf("ID: %v, Name: %v", location.ID, location.Name)
}

func getLocationsFromHttpResponse(resp *http.Response) (*Locations, bool) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, logErrorAndReturnFalse(err)
	}

	return unmarshallXmlResponseIntoLocations(body)
}

func unmarshallXmlResponseIntoLocations(xmldata []byte) (*Locations, bool) {
	locations := &Locations{}
	err := xml.Unmarshal(xmldata, locations)
	if err != nil {
		return nil, logErrorAndReturnFalse(err)
	}
	locationLogger.Trace(locations.Locations)
	return locations, true
}

func logErrorAndReturnFalse(err error) bool {
	locationLogger.Error(err)
	return false
}
