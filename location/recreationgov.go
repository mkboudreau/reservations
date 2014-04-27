package location

import "net/http"

const recreationGovServiceUrl string = "http://ridb.recreation.gov/webservices/RIDBServiceNG.cfc?method=getAllFacilities"

type recreationGov struct {
	serviceUrl string
	cachedData *Locations
}

func newRecreationGovService() *recreationGov {
	return &recreationGov{
		serviceUrl: recreationGovServiceUrl,
	}
}

func (service *recreationGov) GetAllLocations() (*Locations, bool) {
	if service.cachedData != nil {
		return service.cachedData, true
	}
	resp, err := http.Get(service.serviceUrl)
	if err != nil {
		return nil, logErrorAndReturnFalse(err)
	}

	defer resp.Body.Close()
	locations, ok := getLocationsFromHttpResponse(resp)
	if ok && locations != nil && len(locations.Locations) > 0 {
		service.cachedData = locations
	}
	return locations, ok
}

func (service *recreationGov) GetLocationById(legacyId string) (*Location, bool) {
	locations, ok := service.GetAllLocations()
	if !ok {
		return nil, ok
	}

	for _, loc := range locations.Locations {
		if loc.LegacyID == legacyId {
			return loc, true
		}
	}

	return nil, false
}
