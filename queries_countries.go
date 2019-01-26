package geonames

import (
	"encoding/xml"
	"net/url"
)

// Country country info struct
type Country struct {
	// ID geoname ID
	ID int `xml:"geonameId"`
	// Name country name
	Name string `xml:"countryName"`
	// Continent country continent code
	Continent string `xml:"continent"`
	// ContinentName country continent name
	ContinentName string `xml:"continentName"`
	// ISOAlpha3 country code in ISO 3166-1 alpha-3
	ISOAlpha3 string `xml:"isoAlpha3"`
	// FIPSCode FIPS country code
	FIPSCode string `xml:"fipsCode"`
	// Population country population
	Population string `xml:"population"`
	// Capital country capital
	Capital string `xml:"capital"`
	// Languages country languages
	Languages StringSlice `xml:"languages"`
	// Currency country currency
	Currency string `xml:"currencyCode"`
	// PostalCodeFormat country postal code format
	PostalCodeFormat string `xml:"postalCodeFormat"`
	// Area country area in square km
	Area float64 `xml:"areaInSqKm"`
	// North bounding box of mainland
	North float64 `xml:"north"`
	// East bounding box of mainland
	East float64 `xml:"east"`
	// South bounding box of mainland
	South float64 `xml:"south"`
	// West bounding box of mainland
	West float64 `xml:"west"`
}

type countryInfoResponse struct {
	response
	XMLName   xml.Name  `xml:"geonames"`
	Countries []Country `xml:"country"`
}

// CountryInfoQuery country info query struct,
// see http://www.geonames.org/export/web-services.html#countryInfo
type CountryInfoQuery struct {
	query     Query
	countries []string
	lang      string
}

// Lang sets response ISO-639-1 language code, default - english
func (q CountryInfoQuery) Lang(lang string) CountryInfoQuery {
	q.lang = lang
	return q
}

// Get executes query
func (q CountryInfoQuery) Get() ([]Country, error) {
	v := url.Values{}

	if len(q.countries) > 0 {
		for _, c := range q.countries {
			v.Add("country", c)
		}
	}

	if q.lang != "" {
		v.Add("lang", q.lang)
	}

	res := countryInfoResponse{}
	if err := q.query.execute("countryInfo", v, &res); err != nil {
		return []Country{}, err
	}

	return res.Countries, res.error()
}
