package main

import (
	"fmt"

	t "github.com/rexsimiloluwah/distance_metrics/text"
)

func main() {
	fmt.Println(t.Levensthein("kitten", "sitting"))
	fmt.Println(t.Levensthein("a", "b"))
}
