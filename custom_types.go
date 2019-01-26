package geonames

import (
	"encoding/xml"
	"strings"
)

// StringSlice comma-separated array of strings
type StringSlice []string

// UnmarshalXML parse strings slice
func (s *StringSlice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	d.DecodeElement(&v, &start)
	*s = strings.Split(v, ",")

	return nil
}
