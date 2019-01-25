package geonames

import "net/url"

type PlaceNameQuery struct {
	query                Query
	placeName            string
	postalCodeStartsWith string
	countries            []string
}

func (q PlaceNameQuery) PostalCodeStartsWith(startsWith string) PlaceNameQuery {
	q.postalCodeStartsWith = startsWith
	return q
}

func (q PlaceNameQuery) Countries(countries ...string) PlaceNameQuery {
	q.countries = countries
	return q
}

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
