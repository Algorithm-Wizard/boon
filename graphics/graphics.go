package graphics

import (
	"image/color"

	"github.com/Algorithm-Wizard/boon/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	Pos  math.Vectori
	Size math.Vectori
	Img  *ebiten.Image
}

func (d *Draw) MoveTo(pos math.Vector) {
	d.Pos.X = int(pos.X * 65536.0)
	d.Pos.Y = int(pos.Y * 65536.0)
}

func (d *Draw) LineTo(dst math.Vector, clr color.Color) {
	bgn := d.Pos
	end := math.Vectori{
		X: int(dst.X * 65536.0),
		Y: int(dst.Y * 65536.0),
	}
	dx := end.X - d.Pos.X
	dy := end.Y - d.Pos.Y
	if math.Absi(dx) > math.Absi(dy) {
		if bgn.X > end.X {
			bgn, end = end, bgn
		}
		ystp := dy / dx
		yi := bgn.Y
		for xi := bgn.X; xi < end.X; xi += 65536 {
			d.Img.Set(xi>>16, yi>>16, clr)
			yi += ystp
		}
	} else if math.Absi(dy) > math.Absi(dx) {
		if bgn.Y > end.Y {
			bgn, end = end, bgn
		}
		xstp := dx / dy
		xi := bgn.X
		for yi := bgn.Y; yi < end.Y; yi += 65536 {
			d.Img.Set(xi>>16, yi>>16, clr)
			xi += xstp
		}
	}
	d.Img.Set(end.X>>16, end.Y>>16, clr)
}
