package main

import (
	"image/color"
	"log"
	rand "math/rand"
	"strconv"

	"github.com/Algorithm-Wizard/boon/graphics"
	"github.com/Algorithm-Wizard/boon/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const lines int = 1000

type Game struct {
	ballp [lines * 2]math.Vector
	ballv [lines * 2]math.Vector
	size  math.Vector
}

func (g *Game) Update() error {
	for i := 0; i < lines*2; i++ {
		g.ballp[i].X += g.ballv[i].X
		g.ballp[i].Y += g.ballv[i].Y
		if g.ballp[i].X >= g.size.X {
			g.ballp[i].X = (g.size.X * 2) - g.ballp[i].X
			g.ballv[i].X = -g.ballv[i].X
		}
		if g.ballp[i].Y >= g.size.Y {
			g.ballp[i].Y = (g.size.Y * 2) - g.ballp[i].Y
			g.ballv[i].Y = -g.ballv[i].Y
		}
		if g.ballp[i].X < 0 {
			g.ballp[i].X = -g.ballp[i].X
			g.ballv[i].X = -g.ballv[i].X
		}
		if g.ballp[i].Y < 0 {
			g.ballp[i].Y = -g.ballp[i].Y
			g.ballv[i].Y = -g.ballv[i].Y
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{16})
	ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.ActualFPS(), 'f', 3, 64))
	var drawer graphics.Draw
	drawer.Img = screen
	for i := 1; i < lines*2; i += 2 {
		drawer.MoveTo(g.ballp[i-1])
		drawer.LineTo(g.ballp[i], color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.size.X), int(g.size.Y)
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	var game Game = Game{}
	for i := 0; i < lines; i++ {
		game.ballp[i].X = rand.Float32() * 799.0
		game.ballp[i].Y = rand.Float32() * 599.0
		game.ballv[i].X = rand.Float32() * 2.5
		game.ballv[i].Y = rand.Float32() * 2.5
	}
	game.size.X = 400
	game.size.Y = 300
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
