package geonames

import "testing"

func TestPostalCodeQuery_Get(t *testing.T) {
	postalCode := "9011"

	_, err := New(getUserName()).FindPostalCode(postalCode).Get()
	if err != nil {
		t.Errorf("Can't find postal code %s, cause %s", postalCode, err.Error())
	}
}

func TestPostalCodeByPlaceNameQuery_Get(t *testing.T) {
	placeName := "Ettelbruck"

	_, err := New(getUserName()).FindPostalCodeByPlaceName(placeName).Get()
	if err != nil {
		t.Errorf("Can't find postal code with name %s, cause %s", placeName, err.Error())
	}
}
