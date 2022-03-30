package gofractal

var (
	Default = NewView(complex(0.5, 1), complex(-2, -1))
)

type View struct {
	bl, tr complex128
}

func NewView(tr, bl complex128) *View {
	return &View{
		tr: tr,
		bl: bl,
	}
}
