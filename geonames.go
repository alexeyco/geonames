package geonames

const Endpoint = "https://secure.geonames.org/"

func New(userName string) Query {
	return Query{
		userName: userName,
	}
}
