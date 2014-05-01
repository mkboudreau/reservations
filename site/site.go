package site

import (
	"time"
	"strings"
)

type SiteAvailabilityCode string

const (
	SiteAvailable   SiteAvailabilityCode = "A"
	SiteUnavailable                      = "W"
	SiteWalkUp                           = "X"
	SiteReserved                         = "R"
)

type SiteDay struct {
	SiteAvailability SiteAvailabilityCode
	Day              time.Time
}
type Site struct {
	SiteNumber int
	SiteDays   []SiteDay
}

func GetSiteAvailabilityCodeForLetter(letterAnyCase string) SiteAvailabilityCode {
	letter := strings.ToUpper(letterAnyCase)
	if letter == (string)(SiteAvailable) {
		return SiteAvailable
	} else if letter == (string)(SiteUnavailable) {
		return SiteUnavailable
	} else if letter == (string)(SiteWalkUp) {
		return SiteWalkUp
	} else if letter == (string)(SiteReserved) {
		return SiteReserved
	} else {
		return SiteUnavailable
	}
}

func IsValidSiteAvailabilityCode(letterAnyCase string) bool {
	letter := strings.ToUpper(letterAnyCase)
	if letter == (string)(SiteAvailable) {
		return true
	} else if letter == (string)(SiteUnavailable) {
		return true
	} else if letter == (string)(SiteWalkUp) {
		return true
	} else if letter == (string)(SiteReserved) {
		return true
	} else {
		return false
	}
}
