package main

import (
	"log"

	"github.com/alexeyco/geonames"
)

func main() {
	places, _ := geonames.New("demo").FindPostalCode("9011").Get()
	log.Println(places)

	places, _ = geonames.New("demo").FindPostalCodeByPlaceName("Ettelbruck").Get()
	log.Println(places)
}
