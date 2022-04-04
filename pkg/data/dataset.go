package data

import (
	"math/rand"
	"time"
)

func NewIntDataset(possibleAnomaliesPercentage float64, steps, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps int) (result []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	diff := maxVisitorsPerTimeSteps - minVisitorsPerTimeSteps
	possibleAnomalies := newPossibilities(steps, possibleAnomaliesPercentage)
	for n := 1; n < steps; n++ {
		randomValue := minVisitorsPerTimeSteps + r1.Intn(diff)
		result = append(result, randomValue)
	}
	return result
}

func newPossibilities(steps int, possibleAnomaliesPercentage float64) (possibilities []int) {
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	size := int(float64(steps) * possibleAnomaliesPercentage / 100)
	if size > steps {
		size = steps
	}
	for i := 0; i < size; i++ {
		possibilities = append(possibilities, r2.Intn(steps-1))
	}
	return
}

func ParseFloat64(dataset []int) (newDataset []float64) {
	for _, d := range dataset {
		f := float64(d)
		newDataset = append(newDataset, f)
	}
	return
}
