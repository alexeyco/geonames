package geonames

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Query struct {
	userName string
}

func (q Query) FindByPostalCode(postalCode string) PostalCodeQuery {
	return PostalCodeQuery{
		query:      q,
		postalCode: postalCode,
		countries:  []string{},
	}
}

func (q Query) FindByPlaceName(placeName string) PlaceNameQuery {
	return PlaceNameQuery{
		query:     q,
		placeName: placeName,
		countries: []string{},
	}
}

func (q Query) execute(e string, v url.Values) ([]Place, error) {
	req, err := q.request(e, v)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []Place{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Place{}, err
	}

	result := response{}
	if err := json.Unmarshal(body, &result); err != nil {
		return []Place{}, err
	}

	if result.Error != nil {
		return []Place{}, errors.New(result.Error.Message)
	}

	return result.Places, nil
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
