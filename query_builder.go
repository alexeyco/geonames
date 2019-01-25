package geonames

type QueryBuilder struct {
	userName string
}

func (b QueryBuilder) FindByPostalCode(postalCode string) PostalCodeQuery {
	return PostalCodeQuery{
		builder:    b,
		postalCode: postalCode,
		countries:  []string{},
	}
}

func (b QueryBuilder) FindByPlaceName(placeName string) PlaceNameQuery {
	return PlaceNameQuery{
		builder:   b,
		placeName: placeName,
		countries: []string{},
	}
}
