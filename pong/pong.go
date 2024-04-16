package main

import (
	"image/color"
	"log"

	"github.com/Algorithm-Wizard/boon/graphics"
	"github.com/Algorithm-Wizard/boon/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ballp math.Vector
	ballv math.Vector
	size  math.Vector
}

/*type vector struct {
	X float32
	Y float32
}*/

func (g *Game) Update() error {
	g.ballp.X += g.ballv.X
	g.ballp.Y += g.ballv.Y
	if g.ballp.X >= g.size.X {
		g.ballp.X = (g.size.X * 2.0) - g.ballp.X
		g.ballv.X = -g.ballv.X
	}
	if g.ballp.Y >= g.size.Y {
		g.ballp.Y = (g.size.Y * 2.0) - g.ballp.Y
		g.ballv.Y = -g.ballv.Y
	}
	if g.ballp.X < 0 {
		g.ballp.X = -g.ballp.X
		g.ballv.X = -g.ballv.X
	}
	if g.ballp.Y < 0 {
		g.ballp.Y = -g.ballp.Y
		g.ballv.Y = -g.ballv.Y
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{16})
	var drawer graphics.Draw
	drawer.Img = screen
	var mid math.Vector
	mid.X = float32(screen.Bounds().Size().X / 2)
	mid.Y = float32(screen.Bounds().Size().Y / 2)
	drawer.MoveTo(mid)
	drawer.LineTo(g.ballp, color.White)
	screen.Set(int(g.ballp.X), int(g.ballp.Y), color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.size.X), int(g.size.Y)
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	var game Game = Game{}
	game.ballp.X = 1
	game.ballp.Y = 1
	game.ballv.X = 1.25
	game.ballv.Y = .75
	game.size.X = 800
	game.size.Y = 600
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
