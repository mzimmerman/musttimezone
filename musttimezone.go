package musttimezone

import (
	"fmt"
	"log"
	"time"
)

var timezones []*time.Location

func init() {
	tzstrings := []string{
		"US/Central",
		"US/Eastern",
	}
	for _, t := range tzstrings {
		loc, err := time.LoadLocation(t)
		if err != nil {
			log.Fatalf("Error parsing timezone location - %v", err)
		}
		timezones = append(timezones, loc)
	}
}

func Parse(format, value string) (time.Time, error) {
	found := false
	var date time.Time
	var err error
	for _, loc := range timezones {
		date, err = time.ParseInLocation(format, value, loc)
		if err != nil {
			break
		}
		if date.Location().String() == loc.String() {
			// success!
			found = true
			break
		}
	}
	if err != nil {
		return date, err
	}
	if !found {
		return date, fmt.Errorf("Unable to find timezone for date - %s", value)
	}
	return date, nil
}
