package main

import (
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
)

type Enemie struct {
  width int
  height int
  positionX float64
  positionY float64
  color color.RGBA
}

func NewEnemie() []Enemie {
  // add random color
  return []Enemie{}
}

func (e Enemie) draw(screen *ebiten.Image) {
  img := ebiten.NewImage(e.width, e.height)
  img.Fill(e.color)
  geom := ebiten.GeoM{}
  geom.Translate(e.positionX, e.positionY)
  screen.DrawImage(img, &ebiten.DrawImageOptions{
    GeoM: geom,
  })
}
