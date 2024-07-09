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

func NewEnemies() []*Enemie {
  // add random color
  var enemies []*Enemie
  for i := range 3 {
    enemie := &Enemie{
      width: 10,
      height: 10,
      positionX: float64(i * 30),
      positionY: float64(i * 60),
      color: colors.Orange,
    }
    enemies = append(enemies, enemie)
  }
  return enemies
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
