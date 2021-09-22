package vectors

type Vector []float64

func Euclidean() {

}

func Manhattan() {

}

func Minkowski() {

}

func Cosine() {

}

func NormalizedEuclidean() {

}

func Chebyshev() {

}

func (v Vector) Length() int {
	return len(v)
}

func (v Vector) At(index int) float64,error {
	if(index > v.Length() - 1){
		return 0, "Index out of range"
	}

	return v[index]
}
