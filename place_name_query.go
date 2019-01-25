package geonames

import "net/url"

// See http://www.geonames.org/export/web-services.html#postalCodeSearch

// PlaceNameQuery list of postal codes and places for the place name
type PlaceNameQuery struct {
	query                Query
	placeName            string
	postalCodeStartsWith string
	countries            []string
}

// PostalCodeStartsWith sets the first characters of a postal code
func (q PlaceNameQuery) PostalCodeStartsWith(startsWith string) PlaceNameQuery {
	q.postalCodeStartsWith = startsWith
	return q
}

// Countries restricts searching by specified countries
func (q PlaceNameQuery) Countries(countries ...string) PlaceNameQuery {
	q.countries = countries
	return q
}

// Get executes query
func (q PlaceNameQuery) Get() ([]Place, error) {
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

	return q.query.execute("postalCodeSearchJSON", v)
}
