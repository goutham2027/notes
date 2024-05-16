package main

import (
	"fmt"
	"testing"
	"time"
)

type testPair struct {
	date        time.Time
	want        bool
	description string
}

func TestIsDateOlderThanXYears(t *testing.T) {
	// table driven testing with different dates
	// test cases
	// 1. date is older than x years
	// 2. date is not older than x years
	// 3. date is today
	// 4. date is in the future
	// 5. date is in the past
	// 6. date is in the past but not older than x years
	// 7. date is in the future but not older than x years
	// 8. date is in the future but older than x years
	xYears := int64(10)
	today := time.Now()
	futureDateNotOlderThanXYears := time.Now().AddDate(1, 0, 0)
	futureDateOlderThanXYears := time.Now().AddDate(11, 0, 0)
	pastDateNotOlderThanXYears := time.Now().AddDate(-1, 0, 0)
	pastDateOlderThanXYears := time.Now().AddDate(-11, 0, 0)

	var testPairList []testPair
	testPairList = append(testPairList, testPair{time.Time{}, true, fmt.Sprintf("Zero time: %v", time.Time{})})
	testPairList = append(testPairList, testPair{today, false, fmt.Sprintf("today: %v", today)})
	testPairList = append(testPairList, testPair{futureDateNotOlderThanXYears, false, fmt.Sprintf("futureDateNotOlderThanXYears: %v", futureDateNotOlderThanXYears)})
	testPairList = append(testPairList, testPair{futureDateOlderThanXYears, false, fmt.Sprintf("futureDateOlderThanXYears: %v", futureDateOlderThanXYears)})
	testPairList = append(testPairList, testPair{pastDateNotOlderThanXYears, false, fmt.Sprintf("pastDateNotOlderThanXYears: %v", pastDateNotOlderThanXYears)})
	testPairList = append(testPairList, testPair{pastDateOlderThanXYears, true, fmt.Sprintf("pastDateOlderThanXYears: %v", pastDateOlderThanXYears)})

	for _, pair := range testPairList {
		// fmt.Println(pair.description)
		got := IsDateOlderThanXYears(&pair.date, xYears)
		if got != pair.want {
			t.Errorf("%s: got %v, want %v", pair.description, got, pair.want)
		}
	}
}
