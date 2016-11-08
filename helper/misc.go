package helper

import "math"

func ObtenerDistancia(lat1, lon1, lat2, lon2 float64) float64 {
	var r float64 = 6378137
	dLat := rad(lat2 - lat1)
	dLon := rad(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(rad(lat1))*math.Cos(rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := r * c
	return d
}

func rad(x float64) float64 {
	return x * math.Pi / 180
}
