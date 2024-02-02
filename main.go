package main

import "fmt"

func main() {

	
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["france"] = "paris"
	countryCapitalMap["indonesia"] = "jakarta"
	countryCapitalMap["italy"] = "rome"
	countryCapitalMap["indonesia"] = "palembang"
	

	for country := range countryCapitalMap {
		 fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
		 capital, ok := countryCapitalMap["indonesia"]
		 if (ok) {
			fmt.Println("capital of indonesia is", capital)
		 }else {
			fmt.Println("capital of indonesia is not present")
		 }
	}
	
