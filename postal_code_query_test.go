package geonames

import "testing"

func TestPostalCodeQuery_Get(t *testing.T) {
	postalCode := "9011"

	_, err := New(getUserName()).FindByPostalCode(postalCode).Get()
	if err != nil {
		t.Errorf("Can't find postal code %s, cause %s", postalCode, err.Error())
	}
}
