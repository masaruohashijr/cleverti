package data

import (
	"cleverti/pkg/domain"
	"math/rand"
	"time"
)

func NewDataset(siteId domain.Site,
	possibleAnomaliesPercentage float64,
	periodLength, timeStep, minVisitorsPerTimeSteps, maxVisitorsPerTimeSteps int) (ds domain.Dataset) {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	diff := maxVisitorsPerTimeSteps - minVisitorsPerTimeSteps
	possibleAnomalies := newPossibilities(periodLength, possibleAnomaliesPercentage)
	var results []float64
	for n := 1; n < periodLength; n = n + timeStep {
		randomValue := -1
		if possibleAnomalies.Contains(n) {
			randomValue = r1.Intn(3 * maxVisitorsPerTimeSteps)
		} else {
			randomValue = minVisitorsPerTimeSteps + r1.Intn(diff)
		}
		results = append(results, float64(randomValue))
	}
	ds = domain.Dataset{
		SiteId:                  string(siteId),
		TimeAgo:                 periodLength,
		TimeStep:                timeStep,
		OutliersDetectionMethod: "3-sigmas",
		MetricesList:            []string{"Visits"},
		Results:                 results,
	}
	return
}

// steps = 30 , percentage = 5% = #6
func newPossibilities(steps int, possibleAnomaliesPercentage float64) (possibilities *set) {
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	size := int(float64(steps) * possibleAnomaliesPercentage / 100)
	if size > steps {
		size = steps
	}
	possibilities = NewSet()
	for i := 0; i < size; i++ {
		possibilities.Add(r2.Intn(steps - 1))
	}
	return
}
