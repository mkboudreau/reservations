package config

import (
	"encoding/json"
	"fmt"
	"github.com/mkboudreau/loggo"
	"strings"
)

var websiteLogger *loggo.LevelLogger = loggo.DefaultLevelLogger()

const (
	WebsiteAssetName string = "assets/websites.json"
)

type Websites struct {
	Websites []*Website  `json:"websites"`
}
type Website struct {
	Name	string	`json:"name"`
	URL           string                    `json:"url"`
	NumSitesPerRequest int	`json:"numsites_request"`
	ParamsStatic  map[string]string         `json:"static_params"`
	ParamsDynamic ExpectedReplacementParams `json:"dynamic_params"`
}

type WebsiteParamValues ExpectedReplacementParams
type ExpectedReplacementParams struct {
	SiteParam       string `json:"site"`
	StartIndexParam string `json:"start_index"`
	StartDateParam  string `json:"start_date"`
}



func NewDefaultWebsites() (*Websites, bool)  {
	websites, err := Asset(WebsiteAssetName)
	if len(websites) == 0 || err != nil {
		logger.Fatal("Could not find asset:", WebsiteAssetName, ", error:", err)
	}

	return NewWebsitesFromJson(string(websites[:]))
}

func NewWebsitesFromFile(file string) (*Websites, bool)  {
	filedata, ok := readFile(file)
	if !ok {
		return nil, ok
	} else {
		return unmarshalWebsiteJson(filedata)
	}
}

func NewWebsitesFromJson(jsondata string) (*Websites, bool)  {
	return unmarshalWebsiteJson([]byte(jsondata))
}

func (w *Websites) FindWebsiteByKey(key string) *Website {
	for _, website := range w.Websites {
		if website.Name == key {
			return website
		}
	}
	return nil
}

func unmarshalWebsiteJson(jsondata []byte) (*Websites, bool) {
	websites := &Websites{}
	err := json.Unmarshal(jsondata, websites)
	if err != nil {
		return nil, logErrorAndReturnFalse(err)
	}
	websiteLogger.Trace(websites)
	return websites, true
}


func (w *Website) BuildURL(params *WebsiteParamValues) string {
	result := w.URL
	result = addParam(result, buildParams(w.ParamsStatic))
	result = addParam(result, w.ParamsDynamic.QueryStringUsingValues(params))
	return result
}

func (params *ExpectedReplacementParams) QueryStringUsingValues(values *WebsiteParamValues) string {
	result := ""
	addParamKeyValue(result, params.SiteParam, values.SiteParam)
	addParamKeyValue(result, params.StartIndexParam, values.StartIndexParam)
	addParamKeyValue(result, params.StartDateParam, values.StartDateParam)
	return result
}

func buildParams(params map[string]string) string {
	result := ""
	for key, value := range params {
		result = addParamKeyValue(result, key, value)
	}
	return result
}
func addParamKeyValue(current string, key string, value string) string {
	return addParam(current, buildParam(key, value))
}
func addParam(current string, newstr string) string {
	actualLastIndex := len(current) - 1
	questionMark := strings.LastIndex(current, "?")
	ampersand := strings.LastIndex(current, "&")

	if actualLastIndex == questionMark {
		return fmt.Sprintf("%v%v",current, newstr)
	} else if actualLastIndex == ampersand {
		return fmt.Sprintf("%v%v",current, newstr)
	} else {
		return fmt.Sprintf("%v&%v",current, newstr)
	}
}
func buildParam(key string, value string) string {
	return fmt.Sprintf("%v=%v",key,value)
}

func (w *Websites) String() string {
	return fmt.Sprintf("Websites %v", w.Websites)
}
func (w *Website) String() string {
	return fmt.Sprintf("Website at %v", w.URL)
}
/*
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
*/

// http://www.recreation.gov/campsiteCalendar.do?page=calendar&contractCode=NRSO&parkId=70925&calarvdate=06/19/2014&sitepage=true&startIdx=0
// http://www.recreation.gov/campsiteCalendar.do?page=calendar&contractCode=NRSO&parkId=70925&calarvdate=06/19/2014&sitepage=true&startIdx=225
