package geonames

import (
	"strings"
	"testing"
)

func TestCountryCodeQuery_Get(t *testing.T) {
	countryCode, err := New(getUserName()).CountryCode(47.03, 10.2).Get()
	if err != nil {
		t.Errorf("Can't recieve country code, cause %s", err.Error())
	}

	if countryCode != "AT" {
		t.Errorf("Must be \"%s\", \"%s\" given", "AT", countryCode)
	}
}

func TestCountryInfoQuery_Get(t *testing.T) {
	countryCodes := []string{"fr", "de", "it", "es"}

	countries, err := New(getUserName()).CountryInfo(countryCodes...).Get()
	if err != nil {
		t.Errorf("Can't recieve country info (for %s), cause %s", strings.Join(countryCodes, ","), err.Error())
	}

	mustLen := len(countryCodes)
	receivedLen := len(countries)

	if receivedLen != mustLen {
		t.Errorf("Must be %d countries, %d given", mustLen, receivedLen)
	}
}
