package main

// import (
// 	"fmt"
// 	"net/url"
// )

// func main() {
// 	// cookieId := "d17da5fa-ea57-ba5f-ac70-ec5cfa38dd23#1696809600000"

// 	// decodedValue, err := url.QueryUnescape(cookieId)
// 	// if err != nil {
// 	// 	fmt.Errorf("Err while query unescaping", err)
// 	// }
// 	// uuidRE := "([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})"
// 	// dateRE := "(\\d+)?"
// 	// cookiePatternRE := regexp.MustCompile(uuidRE + "#?" + dateRE)
// 	// matches := cookiePatternRE.FindStringSubmatch(decodedValue)
// 	// profileId, err := uuid.FromString(matches[1])
// 	// if err != nil {
// 	// 	fmt.Errorf("Err extracting profileID", err)
// 	// }

// 	// for _, m := range matches {
// 	// 	fmt.Println(m)
// 	// }
// 	// fmt.Println(profileId)

// 	// unixTime, err := strconv.ParseInt("1696809600000", 10, 64)
// 	// if err != nil {
// 	// 	fmt.Errorf("unable to parse timestamp")
// 	// }
// 	// fmt.Println(unixTime)

// 	u, err := url.Parse(reqUrl)
// 	if err != nil {
// 		fmt.Errorf("unable to parse url")
// 	}

// 	queryParams := u.Query().Get("")
// 	fmt.Println(queryParams)
// }
