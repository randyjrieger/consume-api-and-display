package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-population-check/models"
	"io/ioutil"
	"net/http"
)

func main() {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"country": "canada",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population", "application/json", responseBody)
	//Handle Error
	if err != nil {
		fmt.Println("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//sb := string(body)
	//	fmt.Printf(sb)

	parseResponse(body)
}

func parseResponse(response []byte) {
	var stats models.Stats
	json.Unmarshal([]byte(response), &stats)
	//fmt.Printf("%v", stats)

	fmt.Println("Country: ", stats.Data.Country)

	for d := range stats.Data.PopulationCounts {
		fmt.Println("Year:", stats.Data.PopulationCounts[d].Year, " Population: ", stats.Data.PopulationCounts[d].Value)
	}
}
