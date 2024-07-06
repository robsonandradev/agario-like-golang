package main

import (
  "log"
	//"image"
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

type Player struct {
  width int
  height int
  positionX float64
  positionY float64
}

type Game struct {
  player *Player
}

func (g *Game) Update() error {
  if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
    g.player.positionX += 2
  }
  if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
    g.player.positionX -= 2
  }
  if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
    g.player.positionY += 2
  }
  if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
    g.player.positionY -= 2
  }
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  screen.Fill(color.White)
  img := ebiten.NewImage(g.player.width, g.player.height)
  img.Fill(colors.DarkBlue)
  geom := ebiten.GeoM{}
  geom.Translate(g.player.positionX, g.player.positionY)
  screen.DrawImage(img, &ebiten.DrawImageOptions{
    GeoM: geom,
  })
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
  player := &Player{
    width: 40,
    height: 40,
    positionX: 40,
    positionY: 40,
  }
	g := &Game{
    player: player,
  }
  ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Foodie")
  //ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
  ebiten.SetWindowPosition(1300, 0)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
