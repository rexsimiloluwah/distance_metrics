package set

import (
	"reflect"
	"testing"
)

var floatDifferenceThresh float64 = 1e-4

func TestContains(t *testing.T) {
	var tests = []struct {
		s        *set
		el       float64
		expected bool
	}{
		{NewSet([]float64{0, 1, 2, 11, 3, 7, 4, 12, 5}), 2, true},
		{NewSet([]float64{-3.2, -4.5, -5.2, 0, 4.1, 100.0}), -3.2, true},
		{NewSet([]float64{1.2e4, -1.34e5, -3.5e2, -3.4e12}), 12000, true},
		{NewSet([]float64{-3.2, -4.5, -5.2, 0, 4.1, 100.0}), 3.2, false},
	}

	for _, test := range tests {
		if output := test.s.Contains(test.el); output != test.expected {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		s1       *set
		s2       *set
		expected *set
	}{
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5})},
		{NewSet([]float64{}), NewSet([]float64{}), NewSet([]float64{})},
		{NewSet([]float64{0, 1, 2, 3, 4, 5}), NewSet([]float64{5, 6, 7, 8, 9, 10}), NewSet([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
		{NewSet([]float64{-1.1, -2.2, -3.5, -4.2}), NewSet([]float64{-4.23, -1.15, -1.1, 1.1, 1, 3.5}), NewSet([]float64{-1.1, -2.2, -3.5, -4.2, -4.23, -1.15, 1, 1.1, 3.5})},
	}

	for _, test := range tests {
		if output := test.s1.Union(test.s2); !reflect.DeepEqual(output, test.expected) {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestIntersection(t *testing.T) {
	var tests = []struct {
		s1       *set
		s2       *set
		expected *set
	}{
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5})},
		{NewSet([]float64{0, 1, 2, 3, 4, 5}), NewSet([]float64{5, 6, 7, 8, 9, 10}), NewSet([]float64{5})},
		{NewSet([]float64{}), NewSet([]float64{}), NewSet([]float64{})},
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{6, 7, 8, 9, 10}), NewSet([]float64{})},
		{NewSet([]float64{-1.1, -2.2, -3.5, -4.2}), NewSet([]float64{-4.23, -1.15, -1.1, 1.1, 1, 3.5}), NewSet([]float64{-1.1})},
	}

	for _, test := range tests {
		if output := test.s1.Intersection(test.s2); !reflect.DeepEqual(output, test.expected) {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestDiff(t *testing.T) {
	var tests = []struct {
		s1       *set
		s2       *set
		expected *set
	}{
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{})},
		{NewSet([]float64{0, 1, 2, 3, 4, 5}), NewSet([]float64{5, 6, 7, 8, 9, 10}), NewSet([]float64{0, 1, 2, 3, 4})},
		{NewSet([]float64{}), NewSet([]float64{}), NewSet([]float64{})},
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{6, 7, 8, 9, 10}), NewSet([]float64{1, 2, 3, 4, 5})},
		{NewSet([]float64{-1.1, -2.2, -3.5, -4.2}), NewSet([]float64{-4.23, -1.15, -1.1, 1.1, 1, 3.5}), NewSet([]float64{-2.2, -3.5, -4.2})},
	}

	for _, test := range tests {
		if output := test.s1.Diff(test.s2); !reflect.DeepEqual(output, test.expected) {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestJaccard(t *testing.T) {
	var tests = []struct {
		s1       *set
		s2       *set
		expected float64
	}{
		{NewSet([]float64{1.1, 2, 3.1, 4.1, 5.1}), NewSet([]float64{2, 3, 4, 89, 2, 3}), 0.1250},
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5}), 1},
		{NewSet([]float64{}), NewSet([]float64{}), 0},
		{NewSet([]float64{0, 1, 2, 5, 6}), NewSet([]float64{0, 2, 3, 4, 5, 7, 9}), 0.3333},
	}

	for _, test := range tests {
		if output := Jaccard(test.s1, test.s2); output-test.expected > floatDifferenceThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestSorensen(t *testing.T) {
	var tests = []struct {
		s1       *set
		s2       *set
		expected float64
	}{
		{NewSet([]float64{1.1, 2, 3.1, 4.1, 5.1}), NewSet([]float64{2, 3, 4, 89, 2, 3}), 0.1250},
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5}), 1},
		{NewSet([]float64{}), NewSet([]float64{}), 0},
		{NewSet([]float64{0, 1, 2, 5, 6}), NewSet([]float64{0, 2, 3, 4, 5, 7, 9}), 0.3333},
	}

	for _, test := range tests {
		if output := Jaccard(test.s1, test.s2); output-test.expected > floatDifferenceThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestTversky(t *testing.T) {
	var tests = []struct {
		s1       *set
		s2       *set
		alpha    float64
		beta     float64
		expected float64
	}{
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1, 2, 3, 4, 5}), 0.5, 0.5, 1.0},
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{6, 7, 8, 9, 10}), 0.5, 0.5, 0.0},
		{NewSet([]float64{1, 2, 3, 4, 5}), NewSet([]float64{1.1, 2, 3.1, 4.1, 5.1}), 0.5, 0.5, 0.2},
		{NewSet([]float64{}), NewSet([]float64{}), 0, 0, 0},
	}

	for _, test := range tests {
		if output := Tversky(test.s1, test.s2, test.alpha, test.beta); output != test.expected {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}
