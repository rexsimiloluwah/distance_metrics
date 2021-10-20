package main

import (
	"fmt"

	g "github.com/rexsimiloluwah/distance_metrics/geo"
)

func main() {
	// Great circle distance
	fmt.Println(g.GreatCircle(32, 32, 32, 32))

	// Haversine distance
	fmt.Println(g.Haversine(32, 32, 32, 32))
}
