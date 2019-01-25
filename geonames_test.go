package geonames

import "os"

func getUserName() string {
	return os.Getenv("GEONAMES_TEST_USERNAME")
}
