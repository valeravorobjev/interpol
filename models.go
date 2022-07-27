package interpol

// Point struct contains geo coordinate
type Point struct {
	Lat float64 // latitude
	Lon float64 // longitude
}

// PointValue struct contains geo coordinate with value
type PointValue struct {
	Point *Point  // geo coordinate
	Value float64 // value
}

// Vertex contains corner points with values
type Vertex struct {
	BottomLeft  *PointValue // bottom left point with value
	TopLeft     *PointValue // top left point with value
	BottomRight *PointValue // bottom right point with value
	TopRight    *PointValue // top right point with value
}
