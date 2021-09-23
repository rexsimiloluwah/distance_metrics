package vector

import (
	"math"
)

type Vector []float64

// Computes the Minkowski distance between two vectors
// Minkowski distance is the generalized metric form of Manhattan and Euclidean distance
func Minkowski(a Vector, b Vector, p int) float64 {

	if a.Length() == 0 || b.Length() == 0 {
		panic("Vectors a and b cannot be empty")
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	var result float64 = 0
	for i := 0; i < a.Length(); i++ {
		result += math.Pow(math.Abs(a.At(i)-b.At(i)), float64(p))
	}
	return math.Pow(result, 1/float64(p))
}

// Computes the Manhattan distance between two vectors
// Also known as city block distance, L1 norm, Minkowski's L1 distance etc.
func Manhattan(a Vector, b Vector) float64 {
	return Minkowski(a, b, 1)
}

// Computes the Euclidean distance between two vectors
func Euclidean(a Vector, b Vector) float64 {
	return Minkowski(a, b, 2)
}

// Computes the Chebyshev distance between two vectors
func Chebyshev(a Vector, b Vector) float64 {
	var result float64 = 0
	if a.Length() == 0 || b.Length() == 0 {
		panic("Vectors a and b cannot be empty")
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	for i := 0; i < a.Length(); i++ {
		d := math.Abs(a[i] - b[i])
		if d > result {
			result = d
		}
	}
	return result
}

// Computes the Squared euclidean distance between two vectors
func SquaredEuclidean(a Vector, b Vector) float64 {
	return math.Pow(Euclidean(a, b), 2)
}

func Covariance(a Vector, b Vector) float64 {
	if a.Length() == 0 || b.Length() == 0 {
		panic("Vectors a and b cannot be empty")
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	var result float64 = 0
	for i := 0; i < a.Length(); i++ {
		result += (a[i] - a.Mean()) * (b[i] - b.Mean())
	}

	return result / float64(a.Length()-1)
}

// Computes the Pearson Correlation between two vectors
// The returned correlation varies between -1 and 1
// 1 indicates a strong positive correlation between the vectors, and -1 indicates a strong negative correlation, 0 indicates no relationship
func PearsonCorrelation(a Vector, b Vector) float64 {
	var correlation float64 = 0
	if a.Length() == 0 || b.Length() == 0 {
		panic("Vectors a and b cannot be empty")
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	correlation = Covariance(a, b) / (a.Stdev() * b.Stdev())
	return correlation
}

// Computes the Cosing similarity distance between two non-zero vectors
// Cosine similarity basically measures the similarity between two vectors based on the cosine of the angle between them
// Useful for finding similarity between documents in Natural language processing
func CosineSimilarity(a Vector, b Vector) float64 {
	if a.Length() == 0 || b.Length() == 0 {
		panic("Vectors a and b cannot be empty")
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	return a.Dot(b) / (a.Magnitude() * b.Magnitude())
}

// Computes the Cosine dissimilarity between two vectors
func CosineDissimilarity(a Vector, b Vector) float64 {
	return 1 - CosineSimilarity(a, b)
}

// Computes the Hamming distance between two boolean vectors
// Hamming distance basically measures the number of bit positions in the two arrays where the corresponding symbols are different
// Useful for error detection and error correction in data transmitted over computer networks
func Hamming(a Vector, b Vector) int {
	if a.Length() == 0 || b.Length() == 0 {
		panic("Vectors a and b cannot be empty")
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	count := 0
	for i := 0; i < a.Length(); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

/* Utility functions */

// Add two vectors
func Add(a Vector, b Vector) Vector {
	if a.Length() == 0 || b.Length() == 0 {
		return Vector{}
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	sum := make(Vector, a.Length())
	for i := 0; i < a.Length(); i++ {
		sum[i] = a[i] + b[i]
	}
	return sum
}

// Subtract two vectors
func Subtract(a Vector, b Vector) Vector {
	if a.Length() == 0 || b.Length() == 0 {
		return Vector{}
	}

	if a.Length() != b.Length() {
		panic("Vectors a and b must be of the same length.")
	}

	diff := make(Vector, a.Length())
	for i := 0; i < a.Length(); i++ {
		diff[i] = a[i] - b[i]
	}
	return diff
}

// Map an anonymous function to the vector
func Map(a Vector, f func(a interface{}) interface{}) Vector {
	result := make(Vector, a.Length())
	for i := 0; i < a.Length(); i++ {
		result[i] = f(a[i]).(float64)
	}
	return result
}

// Computes the length (n) of an n-dimensional vector
func (v Vector) Length() int {
	return len(v)
}

// Computes the element at a particular index in an n-dimensional vector
func (v Vector) At(index int) float64 {
	if index > v.Length()-1 {
		panic("Vector index out of range.")
	}
	return v[index]
}

// Computes the mean of an n-dimensional vector
func (v Vector) Mean() float64 {
	var sum float64 = 0
	for _, value := range v {
		sum += value
	}
	return sum / float64(v.Length())
}

// Computes the Standard deviation of elements in a vector
func (v Vector) Stdev() float64 {
	var result float64 = 0
	for _, value := range v {
		result += math.Pow(value-v.Mean(), 2)
	}
	return math.Sqrt(result / float64(v.Length()-1))
}

// Computes the magnitude of an n-dimensional vector
func (v Vector) Magnitude() float64 {
	var result float64 = 0
	for _, value := range v {
		result += math.Pow(value, 2)
	}
	return math.Sqrt(result)
}

// Computes the Dot product between a vector and another vector of the same length
func (v Vector) Dot(u Vector) float64 {
	var result float64 = 0
	if v.Length() != u.Length() {
		panic("The length of the vectors must be equal.")
	}

	for i := 0; i < v.Length(); i++ {
		result += v[i] * u[i]
	}
	return result
}

// Computes the index of the maximum element in a a vector
func (v Vector) Maxarg() int {
	maxIdx := 0
	for i := 0; i < v.Length(); i++ {
		if v[i] > v[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

// Computes the index of the minimum element in a vector
func (v Vector) Minarg() int {
	minIdx := 0
	for i := 0; i < v.Length(); i++ {
		if v[i] < v[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// Computes the maximum element in a vector
func (v Vector) Max() float64 {
	maxIdx := v.Maxarg()
	return v[maxIdx]
}

// Computes the minimum element in a vector
func (v Vector) Min() float64 {
	minIdx := v.Minarg()
	return v[minIdx]
}
