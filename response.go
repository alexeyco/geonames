package geonames

import (
	"errors"
)

type errorStatus struct {
	Message string `xml:"message,attr"`
	Value   int    `xml:"value,attr"`
}

type response struct {
	Error *errorStatus `xml:"status"`
}

func (r response) error() error {
	if r.Error != nil {
		return errors.New(r.Error.Message)
	}

	return nil
}
