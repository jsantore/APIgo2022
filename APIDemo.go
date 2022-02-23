package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("What University should we search for:")
	var univName string
	fmt.Scan(&univName)
	apiURL := fmt.Sprintf("http://universities.hipolabs.com/search?name=%s", univName)
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error getting internet response.....\nCowardly quitting.....\n")
		os.Exit(-1)
	}
	defer response.Body.Close()
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return
	}
	universities := make([]UniversityResponse, 2)
	if err = json.Unmarshal(bodyData, &universities); err != nil {
		fmt.Println("Error - could not translate json data to struct properly")
	}
	for _, univData := range universities {
		univData.prettyPrint()
	}
}
