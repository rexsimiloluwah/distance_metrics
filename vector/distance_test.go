package vector

import (
	"math"
	"reflect"
	"testing"
)

var floatDifferenceThresh float64 = 1e-4

func TestAdd(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected Vector
	}{
		{Vector{1, 2, 3, 4, 5}, Vector{0, 0, 0, 0, 0}, Vector{1, 2, 3, 4, 5}},
		{Vector{}, Vector{}, Vector{}},
		{Vector{-1, -2, -3, 4.5, 7, 12}, Vector{-2, -5.2, 5, 2, 12, 0}, Vector{-3, -7.2, 2, 6.5, 19, 12}},
	}

	for _, test := range tests {
		if output := Add(test.v1, test.v2); !reflect.DeepEqual(test.expected, output) {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected Vector
	}{
		{Vector{1, 2, 3, 4, 5}, Vector{0, 0, 0, 0, 0}, Vector{1, 2, 3, 4, 5}},
		{Vector{}, Vector{}, Vector{}},
		{Vector{-1, -2, -3, 4.5, 7, 12}, Vector{-2, -5.2, 5, 2, 12, 0}, Vector{1, 3.2, -8, 2.5, -5, 12}},
	}

	for _, test := range tests {
		if output := Subtract(test.v1, test.v2); !reflect.DeepEqual(test.expected, output) {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestMaxarg(t *testing.T) {
	var tests = []struct {
		v1       Vector
		expected int
	}{
		{Vector{1, 2, 3, 4, 5}, 4},
		{Vector{1e4, -2, -3, 4.5, 12.12, 12}, 0},
		{Vector{-1, -2, -0.2, -4, -5}, 2},
	}

	for _, test := range tests {
		if output := test.v1.Maxarg(); test.expected != output {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestMinarg(t *testing.T) {
	var tests = []struct {
		v1       Vector
		expected int
	}{
		{Vector{1, 2, 3, 4, 5}, 0},
		{Vector{1e4, -2, -3, 4.5, 12.12, 12}, 2},
		{Vector{-1, -2, -0.2, -4, -5}, 4},
	}

	for _, test := range tests {
		if output := test.v1.Minarg(); test.expected != output {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestAt(t *testing.T) {
	var tests = []struct {
		v        Vector
		index    int
		expected float64
	}{
		{Vector{1, 2, 3, 4, 5}, 0, 1},
		{Vector{1e4, -2, -3, 4.5, 12.12, 12}, 1, -2},
		{Vector{-1, -2, -0.2, -4, -5}, 4, -5},
	}

	for _, test := range tests {
		if output := test.v.At(test.index); test.expected != output {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		v        Vector
		expected float64
	}{
		{Vector{1, 2, 3, 4, 5}, 5},
		{Vector{1e4, -2, -3, 4.5, 12.12, 12}, 1e4},
		{Vector{-1, -2, -0.2, -4, -5}, -0.2},
	}

	for _, test := range tests {
		if output := test.v.Max(); test.expected != output {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		v1       Vector
		expected float64
	}{
		{Vector{1, 2, 3, 4, 5}, 1},
		{Vector{1e4, -2, -3, 4.5, 12.12, 12}, -3},
		{Vector{-1, -2, -0.2, -4, -5}, -5},
	}

	for _, test := range tests {
		if output := test.v1.Min(); test.expected != output {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

// // Minkowski is a generalization of Manhattan and Euclidean
func TestMinkowski(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		p        int
		expected float64
	}{
		{Vector{10, 2, 4, -1, 0, 9, 1}, Vector{14, 7, 11, 5, 2, 2, 18}, 4, 17.3452},
		{Vector{1, 2, 3, 4, 5}, Vector{1, 2, 3, 4, 5}, 1, 0},
		{Vector{22, 1, 42, 10}, Vector{20, 0, 36, 8}, 2, 6.7082},
	}

	for _, test := range tests {
		if output := Minkowski(test.v1, test.v2, test.p); math.Abs(output-test.expected) > floatDifferenceThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestManhattan(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected float64
	}{
		{Vector{2, 4, 4, 6}, Vector{5, 5, 7, 8}, 9},
		{Vector{1, 2}, Vector{-1, 4}, 4},
	}

	for _, test := range tests {
		if output := Manhattan(test.v1, test.v2); math.Abs(output-test.expected) > floatDifferenceThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestEuclidean(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected float64
	}{
		{Vector{3, 2}, Vector{4, 1}, math.Sqrt(2)},
	}

	for _, test := range tests {
		if output := Euclidean(test.v1, test.v2); math.Abs(output-test.expected) > floatDifferenceThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestHamming(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected int
	}{
		{Vector{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Vector{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{Vector{0, 1, 1, 1, 0}, Vector{1, 0, 1, 0, 1}, 4},
	}

	for _, test := range tests {
		if output := Hamming(test.v1, test.v2); output != test.expected {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestCosineSimilarity(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected float64
	}{
		{Vector{3, 2, 0, 5}, Vector{1, 0, 0, 1}, 8 / (math.Sqrt(38) * math.Sqrt(2))},
	}

	for _, test := range tests {
		if output := CosineSimilarity(test.v1, test.v2); output != test.expected {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestChebyshev(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected float64
	}{
		{Vector{1, 2, 3, 4, 5, 6, 7, 8}, Vector{2, 4, 6, 8, 10, 20, 11, 16}, 14},
		{Vector{1, 2, 3, 4, 5}, Vector{1, 2, 3, 4, 5}, 0},
	}

	for _, test := range tests {
		if output := Chebyshev(test.v1, test.v2); output != test.expected {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}

func TestPearsonCorrelation(t *testing.T) {
	var tests = []struct {
		v1       Vector
		v2       Vector
		expected float64
	}{
		{Vector{43, 21, 25, 42, 57, 59}, Vector{99, 65, 79, 75, 87, 81}, 0.5298},
		{Vector{1, 2, 3, 4, 5}, Vector{1, 2, 3, 4, 5}, 1},
	}

	for _, test := range tests {
		if output := PearsonCorrelation(test.v1, test.v2); math.Abs(test.expected-output) > floatDifferenceThresh {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}
