package geonames

type Place struct {
	Country    string  `json:"countryCode"`
	PostalCode string  `json:"postalCode"`
	Name       string  `json:"placeName"`
	AdminName1 string  `json:"adminName1"`
	AdminName2 string  `json:"adminName2"`
	AdminCode1 string  `json:"adminCode1"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lng"`
}

type status struct {
	Message string `json:"Message"`
	Value   int    `json:"Value"`
}

type response struct {
	Places []Place `json:"postalCodes,omitempty"`
	Error  *status `json:"status,omitempty"`
}
