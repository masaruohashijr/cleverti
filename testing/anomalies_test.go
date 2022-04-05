package main

import (
	"cleverti/pkg/service"
	"testing"
)

// test function
func TestAnomaliesDetector(t *testing.T) {
	dataset := []float64{9.7, 9.9, 10.5, 10.1, 9.6, 9.7, 9.4, 9.8, 10.0, 9.6, 7, 21}
	outliers, _ := service.Detect(dataset)
	if outliers[0] != 21.0 {
		t.Errorf("Expected outlier(%f) is not same as"+
			" actual value (%f)", outliers[0], 21.0)
	}
}
