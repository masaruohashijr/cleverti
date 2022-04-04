package main

import (
	"fmt"
	"math"
)

func main() {
	dataset := []float64{9.7, 9.9, 10.5, 10.1, 9.6, 9.7, 9.4, 9.8, 10.0, 9.6, 7, 11}
	avg := average(dataset)
	fmt.Printf("Avg %v\n", avg)
	newDataset := NewDataset(dataset, avg)
	fmt.Printf("New Dataset %v\n", newDataset)
	variance := average(newDataset)
	fmt.Printf("Difference %v\n", variance)
	sigma := math.Sqrt(variance)
	fmt.Printf("Sigma %v\n", sigma)
	sigma3Times := 3 * sigma
	fmt.Printf("3-Sigmas %v\n", sigma3Times)
	highEnd := avg + sigma3Times
	fmt.Printf("High End %v\n", highEnd)
	lowEnd := avg - sigma3Times
	fmt.Printf("Low End %v\n", lowEnd)
}

func average(dataset []float64) (avg float64) {
	var sum float64
	var len int = len(dataset)
	for i := 0; i < len; i++ {
		sum += dataset[i]
	}
	avg = (sum / float64(len))
	return
}

func NewDataset(dataset []float64, avg float64) (newdataset []float64) {
	var len int = len(dataset)
	for i := 0; i < len; i++ {
		ele := dataset[i] - avg
		ele = math.Pow(ele, 2)
		newdataset = append(newdataset, ele)
	}
	return
}
