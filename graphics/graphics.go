package graphics

import (
	"image/color"

	"github.com/Algorithm-Wizard/boon/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	pos  math.Vector
	size math.Vector
	img  ebiten.Image
}

func (d *Draw) MoveTo(pos math.Vector) {
	d.pos = pos
}

func (d *Draw) LineTo(dst math.Vector, clr color.Color) {
	bgn := d.pos
	end := dst
	dx := math.Abs(end.X - d.pos.X)
	dy := math.Abs(end.Y - d.pos.Y)
	if dx > dy {
		if bgn.X > end.X {
			bgn, end = end, bgn
		}
		ystp := dy / dx
		yi := bgn.Y
		for xi := bgn.X; xi < end.X; xi += 1.0 {
			d.img.Set(int(xi), int(yi), clr)
			yi += ystp
		}
	} else if dy > dx {
		if bgn.Y > end.Y {
			bgn, end = end, bgn
		}
		xstp := dx / dy
		xi := bgn.X
		for yi := bgn.Y; yi < end; yi += 1.0 {
			d.img.Set(int(xi), int(yi), clr)
			xi += xstp
		}
	}
	d.img.Set(int(d.pos.X), int(d.pos.Y), clr)
	d.img.Set(int(end.X), int(end.Y), clr)
}
