package graphics

import (
	"image/color"

	"github.com/Algorithm-Wizard/boon/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	Pos  math.Vector
	Size math.Vector
	Img  *ebiten.Image
}

func (d *Draw) MoveTo(pos math.Vector) {
	d.Pos = pos
}

func (d *Draw) LineTo(dst math.Vector, clr color.Color) {
	bgn := d.Pos
	end := dst
	dx := end.X - d.Pos.X
	dy := end.Y - d.Pos.Y
	if math.Abs(dx) > math.Abs(dy) {
		if bgn.X > end.X {
			bgn, end = end, bgn
		}
		ystp := dy / dx
		yi := bgn.Y
		for xi := bgn.X; xi < end.X; xi += 1.0 {
			d.Img.Set(int(xi), int(yi), clr)
			yi += ystp
		}
	} else if math.Abs(dy) > math.Abs(dx) {
		if bgn.Y > end.Y {
			bgn, end = end, bgn
		}
		xstp := dx / dy
		xi := bgn.X
		for yi := bgn.Y; yi < end.Y; yi += 1.0 {
			d.Img.Set(int(xi), int(yi), clr)
			xi += xstp
		}
	}
	d.Img.Set(int(d.Pos.X), int(d.Pos.Y), clr)
	d.Img.Set(int(end.X), int(end.Y), clr)
}
