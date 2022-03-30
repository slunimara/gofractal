package gofractal

var (
	PointZero = NewPoint(0, 0)
)

type Point struct {
	x, y uint64
}

// NewPoint returns a new point with the given coordinates.
func NewPoint(x, y uint64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// X returns the x coordinate of the point.
func (p Point) X() uint64 {
	return p.x
}

// SetX sets the x coordinate of the point.
func (p *Point) SetX(x uint64) {
	p.x = x
}

// Y returns the y coordinate of the point.
func (p Point) Y() uint64 {
	return p.y
}

// SetY sets the y coordinate of the point.
func (p *Point) SetY(y uint64) {
	p.y = y
}

// Add adds the given point to the current point.
func (p *Point) Add(other Point) {
	p.x += other.x
	p.y += other.y
}

// AddXY adds the given x and y coordinates to the current point.
func (p *Point) AddXY(x, y uint64) {
	p.Add(NewPoint(x, y))
}
