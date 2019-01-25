package geonames

import "net/url"

type PostalCodeQuery struct {
	query               Query
	postalCode          string
	placeNameStartsWith string
	countries           []string
}

func (q PostalCodeQuery) PlaceNameStartsWith(startsWith string) PostalCodeQuery {
	q.placeNameStartsWith = startsWith
	return q
}

func (q PostalCodeQuery) Countries(countries ...string) PostalCodeQuery {
	q.countries = countries
	return q
}

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
