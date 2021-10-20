// Vector distance metrics
package main

import (
	"fmt"

	v "github.com/rexsimiloluwah/distance_metrics/vector"
)

func main() {
	v1 := v.Vector{1, 2, 3, 4, 5}
	v2 := v.Vector{1, 2, 3, 4, 5}

	// Minkowski
	// Minkowski(v1, v2, p)
	fmt.Println(v.Minkowski(v1, v2, 1))

	// Manhattan
	// Manhattan(v1, v2)
	fmt.Println(v.Manhattan(v1, v2))

	// Euclidean
	// Euclidean(v1,v2)
	fmt.Println(v.Euclidean(v1, v2))

	// Chebyshev
	// Chebyshev(v1,v2)
	fmt.Println(v.Chebyshev(v1, v2))

	// Squared Euclidean
	// SquaredEuclidean(v1,v2)
	fmt.Println(v.SquaredEuclidean(v1, v2))

	// Covariance
	// Covariance(v1,v2)
	fmt.Println(v.Covariance(v1, v2))

	// PearsonCorrelation
	// PearsonCorrelation(v1,v2)
	fmt.Println(v.PearsonCorrelation(v1, v2))

	// CosineSimilarity
	// CosineSimilarity
	fmt.Println(v.CosineSimilarity(v1, v2))

	// Hamming Distance
	fmt.Println(v.Hamming([]float64{0, 0, 1, 0, 0}, []float64{1, 1, 0, 0, 0}))

	// Utility functions
	fmt.Println(v.Add(v1, v2))
	fmt.Println(v.Subtract(v1, v2))
	fmt.Println(v1.Mean())
	fmt.Println(v1.Stdev())
	fmt.Println(v1.Maxarg(), v2.Minarg())
	fmt.Println(v1.Min(), v1.Max())
	fmt.Println(v1.Magnitude())
	fmt.Println(v1.Dot(v2))
}
