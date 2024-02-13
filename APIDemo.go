package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	apiURL := "http://universities.hipolabs.com/search?name=young"
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error getting internet response.....\nCowardly quitting.....\n")
		os.Exit(-1)
	}
	defer response.Body.Close()
	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return
	}
	universities := make([]UniversityResponse, 10, 20)
	if err = json.Unmarshal(bodyData, &universities); err != nil {
		fmt.Println("Error - could not translate json data to struct properly")
	}
	for _, univData := range universities {
		fmt.Println(univData)
	}
}
