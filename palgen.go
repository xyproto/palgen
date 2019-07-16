package palgen

import (
	"image"
	"image/color"
)

func Generate(img image.Image, N int) color.Palette {
	col := color.NRGBA{255, 0, 0, 0}
	pal := color.Palette{col}
	return pal
}
