package geonames

// Endpoint API endpoint
const Endpoint = "https://secure.geonames.org/"

// New creates new API request and returns it
func New(userName string) Query {
	return Query{
		userName: userName,
	}
}

// PostalCode postal code struct
type PostalCode struct {
	// Country ISO3166-2 country name
	Country string `json:"countryCode"`

	// PostalCode postal code
	PostalCode string `json:"postalCode"`

	// Name place name
	Name string `json:"placeName"`

	// AdminName1 admin name
	AdminName1 string `json:"adminName1"`

	// AdminName2 admin name
	AdminName2 string `json:"adminName2"`

	// AdminCode1 admin code
	AdminCode1 string `json:"adminCode1"`

	// Latitude place latitude
	Latitude float64 `json:"lat"`

	// Longitude place longitude
	Longitude float64 `json:"lng"`
}

type status struct {
	Message string `json:"Message"`
	Value   int    `json:"Value"`
}

type response struct {
	Places []PostalCode `json:"postalCodes,omitempty"`
	Error  *status      `json:"status,omitempty"`
}
