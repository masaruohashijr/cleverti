package main

import (
	"bytes"
	"cleverti/pkg/data"
	d "cleverti/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	possibleAnomaliesPercentage := 15.0
	periodLength := 30
	timeStep := 1
	minVisitorsPerTimeSteps := 1000
	maxVisitorsPerTimeSteps := 1500
	ds1 := data.NewDataset(d.Brax, possibleAnomaliesPercentage, periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	ds2 := data.NewDataset(d.Hotjar, possibleAnomaliesPercentage, periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	ds3 := data.NewDataset(d.HubSpot, possibleAnomaliesPercentage, periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	ds4 := data.NewDataset(d.Mandrill, possibleAnomaliesPercentage, periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	ds5 := data.NewDataset(d.Petrus, possibleAnomaliesPercentage, periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	ds6 := data.NewDataset(d.Zendesk, possibleAnomaliesPercentage, periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	datasets := d.Datasets{
		Datasets: []d.Dataset{ds1, ds2, ds3, ds4, ds5, ds6},
	}
	jsonStr, err := json.Marshal(datasets)
	checkErr(err)
	url := "http://localhost:8080/anomalies"
	fmt.Println("URL:>", url)
	println(string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "datasets")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
