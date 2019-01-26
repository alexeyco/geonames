package geonames

// Endpoint API endpoint
const Endpoint = "https://secure.geonames.org/"

// New creates new API request and returns it
func New(userName string) Query {
	return Query{
		userName: userName,
	}
}
