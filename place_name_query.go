package geonames

type PlaceNameQuery struct {
	builder              QueryBuilder
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
	return []Place{}, nil
}
