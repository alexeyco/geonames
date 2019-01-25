package geonames

func New(userName string) QueryBuilder {
	return QueryBuilder{
		userName: userName,
	}
}
