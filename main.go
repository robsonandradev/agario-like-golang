package main

import (
  "log"
	//"image"
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  screen.Fill(color.White)
  img := ebiten.NewImage(50, 50)
  img.Fill(color.Black)
  geom := ebiten.GeoM{}
  x, y := geom.Apply(50, 50)
  geom.Translate(x, y)
  screen.DrawImage(img, &ebiten.DrawImageOptions{
    GeoM: geom,
  })
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cursor (Ebitengine Demo)")
  ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
