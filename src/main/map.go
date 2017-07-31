package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	//traversal map
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	//check exist
	capital, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital of", "United States", "is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
	fmt.Println("------------------- map delete------------------------")
	mapDelete(countryCapitalMap)
}

func mapDelete(countryCapitalMap map[string]string) {

	fmt.Println("the inital map")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	fmt.Println("the map after delete")
	delete(countryCapitalMap, "France")
	delete(countryCapitalMap, "United States")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
