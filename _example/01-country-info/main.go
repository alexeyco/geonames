package main

import (
	"log"

	"github.com/alexeyco/geonames"
)

func main() {
	countries, err := geonames.New("demo").CountryInfo("fr", "de").Get()
	log.Println(countries, err)
}
