package main

import (
	"cleverti/pkg/data"
	"cleverti/pkg/service"
)

func main() {
	// 1
	possibleAnomaliesPercentage := 5.0
	steps := 30
	minVisitorsPerTimeSteps := 1000
	maxVisitorsPerTimeSteps := 1500
	ds := data.NewIntDataset(possibleAnomaliesPercentage, steps, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps)
	fdata := data.ParseFloat64(ds)
	outliers, err := service.Detect(fdata)
	checkErr(err)
	println(outliers)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
