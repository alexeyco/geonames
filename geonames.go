package geonames

func New(userName string) Query {
	return Query{
		userName: userName,
	}
}
