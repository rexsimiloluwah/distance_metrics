package main

import (
	"fmt"

	"github.com/rexsimiloluwah/distance_metrics/set"
	"github.com/rexsimiloluwah/distance_metrics/vector"
)

func main() {
	s1, s2 := set.NewSet([]float64{1, 2, 3, 4, 5}), set.NewSet([]float64{6, 7, 8, 9, 10})
	v1, v2 := vector.Vector{1, 2, 3, 4, 5}, vector.Vector{1, 2, 3, 4, 5}

	fmt.Println(v1, "+", v2, " = ", vector.Add(v1, v2))
	fmt.Println(s1, "U", s2, " = ", s1.Union(s2))
}
