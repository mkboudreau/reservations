package config

import (
	"time"
	"encoding/json"
	"github.com/mkboudreau/loggo"
	"fmt"
)

var monitorLogger *loggo.LevelLogger = loggo.DefaultLevelLogger()

const (
	MonitorAssetName string = "assets/monitors.json"
)

type Monitors struct {
	Monitors map[string]*Monitor `json:monitors`
}

type Monitor struct {
	Name              string      `json:"name"`
	Location          int         `json:"location"`
	WebsiteKey          string         `json:"websiteKey"`
	Sites             []int       `json:"sites"`
	HighPrioritySites []int       `json:"hpsites"`
	HighPriorityDates []*MonitorTime `json:"hpdates"`
	From              *MonitorTime   `json:"from"`
	To                *MonitorTime   `json:"to"`
	Emails            []string    `json:"emails"`
}

type MonitorTime struct {
	Time time.Time
}

func (m *Monitors) String() string {
	return fmt.Sprintf("Monitors %v", m.Monitors)
}
func (m *Monitor) String() string {
	return fmt.Sprintf("Monitor %v for location %v from %v to %v", m.Name, m.Location, m.From, m.To)
}
func (m *MonitorTime) String() string {
	return fmt.Sprintf("%v-%v-%v", m.Time.Year(), m.Time.Month(), m.Time.Day())
}
func (t *MonitorTime) UnmarshalJSON(jsondata []byte) error {
	dateString := string(jsondata[1:len(jsondata)-1]) + " 00:00:00 -0700"

	tm, err := time.Parse("2006-01-02 15:04:05 -0700", dateString)
	if err != nil {
		return err
	}

	t.Time = tm
	return nil
}

func NewDefaultMonitors() (*Monitors, bool)  {
	monitors, err := Asset(MonitorAssetName)
	if len(monitors) == 0 || err != nil {
		logger.Fatal("Could not find asset:", MonitorAssetName, ", error:", err)
	}

	return NewMonitorsFromJson(string(monitors[:]))
}

func NewMonitorsFromFile(file string) (*Monitors, bool)  {
	filedata, ok := readFile(file)
	if !ok {
		return nil, ok
	} else {
		return unmarshalMonitorJson(filedata)
	}
}

func NewMonitorsFromJson(jsondata string) (*Monitors, bool)  {
	return unmarshalMonitorJson([]byte(jsondata))
}


func unmarshalMonitorJson(jsondata []byte) (*Monitors, bool) {
	monitors := &Monitors{}
	err := json.Unmarshal(jsondata, monitors)
	if err != nil {
		return nil, logErrorAndReturnFalse(err)
	}
	monitorLogger.Trace(monitors)
	return monitors, true
}
