package geo

import "math"

var EARTH_RADIUS float64 = 6371

// Compute Great circle distance between two points
// @param {float64} lon1 : Longitude of first point
// @param {float64} lon2 : Longitude of second point
// @param {float64} lat1 : Latitude of first point
// @param {float64} lat2 : Latitude of second point
func GreatCircle(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {
	lon1, lat1, lon2, lat2 = ToRad(lon1), ToRad(lat1), ToRad(lon2), ToRad(lat2)
	if lon1 == lon2 && lat1 == lat2 {
		return 0
	}
	return EARTH_RADIUS * math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon1-lon2))
}

// Compute Haversine distance between two points
// @param {float64} lon1 : Longitude of first point
// @param {float64} lon2 : Longitude of second point
// @param {float64} lat1 : Latitude of first point
// @param {float64} lat2 : Latitude of second point
func Haversine(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {
	lon1, lat1, lon2, lat2 = ToRad(lon1), ToRad(lat1), ToRad(lon2), ToRad(lat2)
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	return 2 * EARTH_RADIUS * math.Asin(math.Sqrt(a))
}

// TODO: Implement Vincenty iterative algorithm

// Convert degress to radians
func ToRad(deg float64) float64 {
	return (math.Pi / 180) * deg
}
