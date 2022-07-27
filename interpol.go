package interpol

import "math"

// Linear is the linear interpolation
func Linear(x1 float64, x2 float64, q1 float64, q2 float64, x float64) float64 {
	r := ((x2-x)/(x2-x1))*q1 + ((x-x1)/(x2-x1))*q2
	return r
}

// Bilinear is the bilinear interpolation
func Bilinear(vertex *Vertex, point *Point) float64 {
	if vertex == nil {
		return math.NaN()
	}

	x := point.Lon
	y := point.Lat

	x1 := vertex.TopLeft.Point.Lon
	y1 := vertex.TopLeft.Point.Lat
	x2 := vertex.BottomRight.Point.Lon
	y2 := vertex.BottomRight.Point.Lat

	q11 := vertex.BottomLeft.Value
	q12 := vertex.TopLeft.Value
	q21 := vertex.BottomRight.Value
	q22 := vertex.TopRight.Value

	r1 := Linear(x1, x2, q11, q21, x)
	r2 := Linear(x1, x2, q12, q22, x)

	p := Linear(y1, y2, r1, r2, y)

	return p

}
