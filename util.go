package main

import (
  "image/color"
)

type Colors struct {
  DarkBlue   color.RGBA
  White      color.RGBA
  Black      color.RGBA
  LightBlue  color.RGBA
  DarkRed    color.RGBA
  LightRed   color.RGBA
  Orange     color.RGBA
  Yellow     color.RGBA
  DarkGreen  color.RGBA
  LightGreen color.RGBA
  Purple     color.RGBA
  Pink       color.RGBA
}

var (
  colors = Colors{
    White: color.RGBA{255, 255, 255, 255},
    Black: color.RGBA{0, 0, 0, 255},
    DarkBlue: color.RGBA{0, 0, 153, 255},
    LightBlue: color.RGBA{51, 153, 255, 255},
    DarkRed: color.RGBA{153, 0, 0, 255},
    LightRed: color.RGBA{255, 51, 51, 255},
    Orange: color.RGBA{255, 128, 0, 255},
    Yellow: color.RGBA{255, 255, 51, 255},
    DarkGreen: color.RGBA{51, 102, 0, 255},
    LightGreen: color.RGBA{51, 255, 51, 255},
    Purple: color.RGBA{102, 0, 204, 255},
    Pink: color.RGBA{204, 0, 204, 255},
  }
)
