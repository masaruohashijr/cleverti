package service

import (
	"fmt"
	"math"
)

func Average(dataset []float64) (avg float64) {
	var sum float64
	var len int = len(dataset)
	for i := 0; i < len; i++ {
		sum += dataset[i]
	}
	avg = (sum / float64(len))
	return
}

func NewInterimDataset(dataset []float64, avg float64) (newInterimDataset []float64) {
	var len int = len(dataset)
	for i := 0; i < len; i++ {
		ele := dataset[i] - avg
		ele = math.Pow(ele, 2)
		newInterimDataset = append(newInterimDataset, ele)
	}
	return
}

func Variance(dataset []float64) (avg float64, variance float64) {
	avg = Average(dataset)
	fmt.Printf("Avg %v\n", avg)
	newDataset := NewInterimDataset(dataset, avg)
	fmt.Printf("New Dataset %v\n", newDataset)
	variance = Average(newDataset)
	fmt.Printf("Variance %v\n", variance)
	return
}

func Detect(dataset []float64) ([]float64, error) {
	avg, variance := Variance(dataset)
	sigma := math.Sqrt(variance)
	fmt.Printf("Sigma %v\n", sigma)
	sigma3Times := 3 * sigma
	fmt.Printf("3-Sigmas %v\n", sigma3Times)
	upperLimit := avg + sigma3Times
	fmt.Printf("Upper Limit %v\n", upperLimit)
	lowerLimit := avg - sigma3Times
	fmt.Printf("Lower Limit %v\n", lowerLimit)
	var outliers []float64
	var err error
	for _, ele := range dataset {
		if ele < lowerLimit || ele > upperLimit {
			outliers = append(outliers, ele)
		}
	}
	if outliers == nil {
		err = fmt.Errorf("No outlier detected %v.\n", outliers)
	}
	return outliers, err
}
