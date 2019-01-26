[![Travis](https://img.shields.io/travis/alexeyco/geonames.svg)](https://travis-ci.org/alexeyco/geonames)
[![Coverage Status](https://coveralls.io/repos/github/alexeyco/geonames/badge.svg?branch=master)](https://coveralls.io/github/alexeyco/geonames?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexeyco/geonames)](https://goreportcard.com/report/github.com/alexeyco/geonames)
[![GoDoc](https://godoc.org/github.com/alexeyco/geonames?status.svg)](https://godoc.org/github.com/alexeyco/geonames)
[![license](https://img.shields.io/github/license/alexeyco/geonames.svg)](https://github.com/alexeyco/geonames)

# Geonames
Rest API client for geonames.org

1. [Countries](#countries)
    1. [Get the ISO country code of any given point](#get-the-iso-country-code-of-any-given-point)
    1. [Get country information by ISO country code](#get-country-information-by-iso-country-code)

## Countries
### Get the ISO country code of any given point

```go
package main

import (
	"log"

	"github.com/alexeyco/geonames"
)

func main() {
	countryCode, _ := geonames.New("demo").CountryCode(47.03, 10.2).
		//Lang("fr").  // response language code in ISO-639-1 (en,de,fr,it,es,...), default = english
		//Radius(100). // buffer in km for closest country in coastal areas
		Get()

	log.Println(countryCode)
}
```

See [documentation](http://www.geonames.org/export/web-services.html#countrycode).

### Get country information by ISO country code

```go
package main

import (
	"log"

	"github.com/alexeyco/geonames"
)

func main() {
	countries, _ := geonames.New("demo").CountryInfo("fr", "de").
		//Lang("fr").  // response language code in ISO-639-1 (en,de,fr,it,es,...), default = english
		Get()

	log.Println(countries)
}
```

See [documentation](http://www.geonames.org/export/web-services.html#countryInfo).

## License

```
MIT License

Copyright (c) 2019 Alexey Popov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
