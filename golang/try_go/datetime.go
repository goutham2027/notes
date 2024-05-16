package main

import (
	"fmt"
	"time"
)

func main() {
	// datetime, _ := time.Parse("2006-01-02 15:04:05 UTC", "2013-03-23 15:04:05 UTC")
	// datetime := time.Now().AddDate(0, 5, 0)
	// datetime := time.Time{}

	// date1_parsed := time.Time(*datetime)
	fmt.Println(datetime)
	fmt.Println(IsDateOlderThanXYears(&datetime, 10))

}

// check if event is older than X years
func IsDateOlderThanXYears(eventDate *time.Time, xYears int64) bool {
	today := time.Now()
	if eventDate.IsZero() {
		return true
	}
	if eventDate.After(today) {
		return false
	}
	// duration := today.Sub(*eventDate)
	// fmt.Println(math.Ceil(duration / time.Hour))
	cutOff := time.Now().AddDate(-int(xYears), 0, 0)
	return eventDate.Before(cutOff)
	// years := float64(duration) / float64(time.Hour) / 24 / 365
	// years := duration / time.Hour / 24 / 365
	// fmt.Println(years)
	// return years > float64(xYears)
	// return int64(years) > xYears

}
