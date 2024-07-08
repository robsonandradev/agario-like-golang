package main

import (
  "log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 600
	screenHeight = 805
)

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
  screen.Fill(colors.White)
  g.player.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
  return screenWidth, screenHeight
}

func main() {
  player := NewPlayer()
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
