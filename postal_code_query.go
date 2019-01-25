package geonames

import "net/url"

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
func (q PostalCodeQuery) Get() ([]Place, error) {
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

	return q.query.execute("postalCodeSearchJSON", v)
}
