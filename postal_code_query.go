package geonames

type PostalCodeQuery struct {
	builder             QueryBuilder
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
	return []Place{}, nil
}