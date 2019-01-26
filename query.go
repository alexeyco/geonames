package geonames

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Query API request structure
type Query struct {
	userName string
}

// CountryCode returns the iso country code of any given point
func (q Query) CountryCode(latitude, longitude float32) CountryCodeQuery {
	return CountryCodeQuery{
		query:     q,
		latitude:  latitude,
		longitude: longitude,
	}
}

// CountryInfo returns country information: capital, population, area in square km, bounding Box of mainland
// (excluding offshore islands); default - all countries
func (q Query) CountryInfo(countries ...string) CountryInfoQuery {
	return CountryInfoQuery{
		query:     q,
		countries: countries,
	}
}

func (q Query) execute(e string, v url.Values, result interface{}) error {
	req, err := q.request(e, v)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}

func (q Query) request(e string, v url.Values) (*http.Request, error) {
	req, err := http.NewRequest("GET", Endpoint+e, nil)
	if err != nil {
		return nil, err
	}

	v.Add("username", q.userName)
	req.URL.RawQuery = v.Encode()

	return req, nil
}
