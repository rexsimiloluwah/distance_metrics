package set

import "fmt"

type set struct {
	data map[float64]bool
	size int
}

var exists bool = true

// Compute the Jaccard Similarity Index between two sets
// Reference: http://en.wikipedia.org/wiki/Jaccard_index
// Jaccard(A,B) = |A n B| / |A u B|
// Similarity index ranges between 0 and 1 -> (0,1)
func Jaccard(s1 *set, s2 *set) float64 {
	if s1.Union(s2).size == 0 {
		return 0
	}
	num := s1.Intersection(s2).size
	den := s1.Union(s2).size
	return float64(num) / float64(den)
}

// Computes the Sorensen-Dice coefficient between two sets
// Sorensen(A,B) = 2|A n B| / (|A| + |B|)
// Reference: https://effectivesoftwaredesign.com/2019/02/27/data-science-set-similarity-metrics/
// Similarity index ranges between 0 and 1 -> (0,1)
func Sorensen(s1 *set, s2 *set) float64 {
	if s1.size == 0 && s2.size == 0 {
		return 0
	}
	num := 2 * (s1.Intersection(s2).size)
	den := s1.size + s2.size
	fmt.Println(num, den)
	return float64(num) / float64(den)
}

// Computes the Tversky index between two sets
// Tversky is an asymmetric similarity measure between two sets, a generalization of the Jaccard and Sorensen coefficient
// Reference: https://en.wikipedia.org/wiki/Tversky_index
// Tversky(A,B) = |A n B| / (|A n B| + alpha*|A - B| + beta*|B - A|) where alpha,beta >= 0
func Tversky(s1 *set, s2 *set, alpha float64, beta float64) float64 {
	if s1.size == 0 && s2.size == 0 {
		return 0.0
	}
	if alpha == 0 && beta == 0 {
		return 1.1
	}

	var num, den float64
	num = float64(s1.Intersection(s2).size)
	den = float64(s1.Intersection(s2).size) + alpha*float64(s1.Diff(s2).size) + beta*float64(s2.Diff(s1).size)
	return num / den
}

// Initialize a new set from an array (Supports only numbers for now)
func NewSet(v []float64) *set {
	s := &set{}
	s.data = make(map[float64]bool)
	for _, value := range v {
		s.data[value] = exists
		s.size++
	}
	return s
}

// Check if the set contains a specific element
func (s *set) Contains(el float64) bool {
	return s.data[el]
}

// Add a new element to the set
func (s *set) Add(el float64) {
	if !s.Contains(el) {
		s.data[el] = exists
		s.size++
	}
}

// Remove an existing element from the set
func (s *set) Remove(el float64) {
	if s.Contains(el) {
		delete(s.data, el)
		s.size--
	} else {
		panic("This element does not exist in the set.")
	}
}

// Convert a set to an array
func (s *set) ToArray() []float64 {
	arr := make([]float64, s.size)
	i := 0
	for el, _ := range *&(s).data {
		if ok := s.Contains(el); ok {
			arr[i] = el
			i++
		}
	}
	return arr
}

// Compute intersection between two sets
func (s1 *set) Intersection(s2 *set) *set {
	if s1.size == 0 || s2.size == 0 {
		return NewSet([]float64{})
	}

	interSet := NewSet([]float64{})
	for el, _ := range *&(s1).data {
		if ok := s2.Contains(el); ok {
			interSet.Add(el)
		}
	}
	return interSet
}

// Compute union between two sets
func (s1 *set) Union(s2 *set) *set {
	if s1.size == 0 {
		return s2
	}

	if s2.size == 0 {
		return s1
	}
	unionSet := NewSet([]float64{})
	for el, _ := range *&(s1).data {
		unionSet.Add(el)
	}
	for el, _ := range *&(s2).data {
		unionSet.Add(el)
	}
	return unionSet
}

// Computes the difference between two sets
// i.e. s1 - s2 returns the elements of s1 which are not elements of s2
func (s1 *set) Diff(s2 *set) *set {
	if s2.size == 0 {
		return s1 // completely disjoint
	}

	diffSet := NewSet([]float64{})
	for el, _ := range *&(s1).data {
		if !s2.Contains(el) {
			diffSet.Add(el)
		}
	}
	return diffSet
}
