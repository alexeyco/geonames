package geonames

import (
	"errors"
	"net/url"
)

// See http://www.geonames.org/export/web-services.html#postalCodeSearch

// PostalCodeQuery list of postal codes and places for the postal code
type PostalCodeQuery struct {
	query               Query
	postalCode          string
	placeNameStartsWith string
	countries           []string
}

// PlaceNameStartsWith sets the first characters of a place name
func (q PostalCodeQuery) PlaceNameStartsWith(startsWith string) PostalCodeQuery {
	q.placeNameStartsWith = startsWith
	return q
}

// Countries restricts searching by specified countries
func (q PostalCodeQuery) Countries(countries ...string) PostalCodeQuery {
	q.countries = countries
	return q
}

// Get executes query
func (q PostalCodeQuery) Get() ([]PostalCode, error) {
	v := url.Values{}
	v.Add("postalcode", q.postalCode)

	if q.placeNameStartsWith != "" {
		v.Add("placename_startsWith", q.placeNameStartsWith)
	}

	if len(q.countries) > 0 {
		for _, c := range q.countries {
			v.Add("country", c)
		}
	}

	res := response{}
	if err := q.query.execute("postalCodeSearchJSON", v, &res); err != nil {
		return []PostalCode{}, err
	}

	if res.Error != nil {
		return []PostalCode{}, errors.New(res.Error.Message)
	}

	return res.Places, nil
}

// PostalCodeByPlaceNameQuery list of postal codes and places for the place name
type PostalCodeByPlaceNameQuery struct {
	query                Query
	placeName            string
	postalCodeStartsWith string
	countries            []string
}

// PostalCodeStartsWith sets the first characters of a postal code
func (q PostalCodeByPlaceNameQuery) PostalCodeStartsWith(startsWith string) PostalCodeByPlaceNameQuery {
	q.postalCodeStartsWith = startsWith
	return q
}

// Countries restricts searching by specified countries
func (q PostalCodeByPlaceNameQuery) Countries(countries ...string) PostalCodeByPlaceNameQuery {
	q.countries = countries
	return q
}

// Get executes query
func (q PostalCodeByPlaceNameQuery) Get() ([]PostalCode, error) {
	v := url.Values{}
	v.Add("placename", q.placeName)

	if q.postalCodeStartsWith != "" {
		v.Add("postalcode_startsWith", q.postalCodeStartsWith)
	}

	if len(q.countries) > 0 {
		for _, c := range q.countries {
			v.Add("country", c)
		}
	}

	res := response{}
	if err := q.query.execute("postalCodeSearchJSON", v, &res); err != nil {
		return []PostalCode{}, err
	}

	if res.Error != nil {
		return []PostalCode{}, errors.New(res.Error.Message)
	}

	return res.Places, nil
}
