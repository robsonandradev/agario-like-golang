package main

import (
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
  width int
  height int
  positionX float64
  positionY float64
  color color.RGBA
}

func NewPlayer() *Player {
  return  &Player{
    width: 20,
    height: 20,
    positionX: 200,
    positionY: 700,
    color: colors.DarkGreen,
  }
}

func (p *Player) draw(screen *ebiten.Image) {
  img := ebiten.NewImage(p.width, p.height)
  img.Fill(p.color)
  geom := ebiten.GeoM{}
  geom.Translate(p.positionX, p.positionY)
  screen.DrawImage(img, &ebiten.DrawImageOptions{
    GeoM: geom,
  })
}
