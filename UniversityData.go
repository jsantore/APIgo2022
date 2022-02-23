package main

import "fmt"

type UniversityResponse struct {
	TwoLetterCountryCode string   `json:"alpha_two_code"`
	Domains              []string `json:"domains"`
	WebPages             []string `json:"web_pages"`
	Country              string   `json:"country"`
	Name                 string   `json:"name"`
	StateProvince        string   `json:"state-province"`
}

func (uni UniversityResponse) prettyPrint() {
	firstOutput := fmt.Sprintf("University:\t%s\nCountry:\t%s\n", uni.Name, uni.Country)
	fmt.Println(firstOutput)
}
