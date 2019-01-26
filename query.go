package geonames

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Query API request structure
type Query struct {
	userName string
}

// FindPostalCode find postal code
func (q Query) FindPostalCode(postalCode string) PostalCodeQuery {
	return PostalCodeQuery{
		query:      q,
		postalCode: postalCode,
		countries:  []string{},
	}
}

// FindPostalCodeByPlaceName find postal code by place name
func (q Query) FindPostalCodeByPlaceName(placeName string) PostalCodeByPlaceNameQuery {
	return PostalCodeByPlaceNameQuery{
		query:     q,
		placeName: placeName,
		countries: []string{},
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

	if err := json.Unmarshal(body, &result); err != nil {
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
