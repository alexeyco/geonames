package geonames

import (
	"strings"
	"testing"
)

func TestCountryInfoQuery_Get(t *testing.T) {
	countryCodes := []string{"fr", "de", "it", "es"}

	countries, err := New(getUserName()).CountryInfo(countryCodes...).Get()
	if err != nil {
		t.Errorf("Can't recieve country info (for %s), cause %s", strings.Join(countryCodes, ","), err.Error())
	}

	mustLen := len(countryCodes)
	receivedLen := len(countries)

	if receivedLen != mustLen {
		t.Errorf("Must be %d countries, %d received", mustLen, receivedLen)
	}
}
