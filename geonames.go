package geonames

import (
	"fmt"
	"strings"
)

// Endpoint API endpoint
const Endpoint = "https://secure.geonames.org/"

// New creates new API request and returns it
func New(userName string) Query {
	return Query{
		userName: userName,
	}
}

func floatToString(v float32) string {
	s := fmt.Sprintf("%.2f", v)
	return strings.TrimRight(s, "0")
}

func intToString(v int) string {
	return fmt.Sprintf("%d", v)
}
