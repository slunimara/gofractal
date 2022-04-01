package gofractal

var (
	Default = NewView(complex(0.5, 1), complex(-2, -1))
)

type View struct {
	// Top right corner
	tr complex128
	// Bottom left corner
	bl complex128
}

func NewView(topRight, bottomLeft complex128) *View {
	return &View{
		tr: topRight,
		bl: bottomLeft,
	}
}
