package main

import (
	"fmt"

	s "github.com/rexsimiloluwah/distance_metrics/set"
)

func main() {
	s1 := s.NewSet([]float64{0, 1, 2, 3, 4, 5})
	s2 := s.NewSet([]float64{5, 6, 7, 8, 9, 10})

	// Jaccard
	fmt.Println(s.Jaccard(s1, s2))

	// Sorensen
	fmt.Println(s.Sorensen(s1, s2))

	// Tversky
	fmt.Println(s.Tversky(s1, s2, 0.5, 0.5))

	// Utility functions
	s1.Add(12)
	fmt.Println(s1.ToArray())

	s1.Remove(12)
	fmt.Println(s1.ToArray())

	// Union
	fmt.Println(s1.Union(s2))

	// Intersection
	fmt.Println(s1.Intersection(s2))

	// Size
	fmt.Println(s1.Size())

	// Difference
	fmt.Println(s1.Diff(s2))
}
