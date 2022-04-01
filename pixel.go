package gofractal

type Pixel struct {
	x, y uint64
}

// NewPixel returns a new pixel with the given coordinates.
func NewPixel(x, y uint64) Pixel {
	return Pixel{
		x: x,
		y: y,
	}
}

// PixelZero returns a new pixel with the coordinates 0, 0.
func PixelZero() Pixel {
	return NewPixel(0, 0)
}

// X returns the x coordinate of the pixel.
func (p Pixel) X() uint64 {
	return p.x
}

// SetX sets the x coordinate of the pixel.
func (p *Pixel) SetX(x uint64) {
	p.x = x
}

// Y returns the y coordinate of the pixel.
func (p Pixel) Y() uint64 {
	return p.y
}

// SetY sets the y coordinate of the pixel.
func (p *Pixel) SetY(y uint64) {
	p.y = y
}

// Add adds the given pixel to the current pixel.
func (p *Pixel) Add(other Pixel) {
	p.x += other.x
	p.y += other.y
}

// AddXY adds the given x and y coordinates to the current pixel.
func (p *Pixel) AddXY(x, y uint64) {
	p.Add(NewPixel(x, y))
}

// AddX adds the given x coordinate to the current pixel.
func (p *Pixel) AddX(x uint64) {
	p.Add(NewPixel(x, 0))
}

// AddY adds the given y coordinate to the current pixel.
func (p *Pixel) AddY(y uint64) {
	p.Add(NewPixel(0, y))
}
