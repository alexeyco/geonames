package geonames

import "testing"

func TestPlaceNameQuery_Get(t *testing.T) {
	placeName := "Ettelbruck"

	_, err := New(getUserName()).FindByPlaceName(placeName).Get()
	if err != nil {
		t.Errorf("Can't find place code %s, cause %s", placeName, err.Error())
	}
}
