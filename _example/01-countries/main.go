package main

import (
	"log"

	"github.com/alexeyco/geonames"
)

func main() {
	c := geonames.New("demo")

	// Get the ISO country code of any given point
	countryCode, _ := c.CountryCode(47.03, 10.2).
		//Lang("fr").  // response language code in ISO-639-1 (en,de,fr,it,es,...), default = english
		//Radius(100). // buffer in km for closest country in coastal areas
		Get()

	log.Println(countryCode)

	// Get country information by ISO country code
	countries, _ := c.CountryInfo("fr", "de").
		//Lang("fr").  // response language code in ISO-639-1 (en,de,fr,it,es,...), default = english
		Get()

	log.Println(countries)
}
