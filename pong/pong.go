package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ballp vector
	ballv vector
	size  vector
}

/*type vector struct {
	x float32
	y float32
}*/

func (g *Game) Update() error {
	g.ballp.x += g.ballv.x
	g.ballp.y += g.ballv.y
	if g.ballp.x >= g.size.x {
		g.ballp.x = (g.size.x * 2.0) - g.ballp.x
		g.ballv.x = -g.ballv.x
	}
	if g.ballp.y >= g.size.y {
		g.ballp.y = (g.size.y * 2.0) - g.ballp.y
		g.ballv.y = -g.ballv.y
	}
	if g.ballp.x < 0 {
		g.ballp.x = -g.ballp.x
		g.ballv.x = -g.ballv.x
	}
	if g.ballp.y < 0 {
		g.ballp.y = -g.ballp.y
		g.ballv.y = -g.ballv.y
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{16})
	screen.Set(int(g.ballp.x), int(g.ballp.y), color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.size.x), int(g.size.y)
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	var game Game = Game{}
	game.ballp.x = 1
	game.ballp.y = 1
	game.ballv.x = 1.25 / 4.0
	game.ballv.y = .75 / 4.0
	game.size.x = 160
	game.size.y = 120
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
