package geo

import (
	"testing"
)

var floatThresh = 1e-2

func TestGreatCircle(t *testing.T) {
	var tests = []struct {
		lon1     float64
		lat1     float64
		lon2     float64
		lat2     float64
		expected float64
	}{
		{32, 32, 32, 32, 0},
		{12, 50.2, 32, 56.1, 1479.08},
	}

	for _, test := range tests {
		if output := GreatCircle(test.lon1, test.lat1, test.lon2, test.lat2); output-test.expected > floatThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestHaversine(t *testing.T) {
	var tests = []struct {
		lon1     float64
		lat1     float64
		lon2     float64
		lat2     float64
		expected float64
	}{
		{32, 32, 32, 32, 0},
		{35.8, 32, 19, 12, 2810.16},
	}

	for _, test := range tests {
		if output := Haversine(test.lon1, test.lat1, test.lon2, test.lat2); output-test.expected > floatThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}

}
