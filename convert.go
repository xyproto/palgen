package palgen

import (
	"fmt"
	"image"
	"image/color"
)

func quickDistance(a, b color.Color) uint {
	c1 := a.(color.RGBA)
	c2 := b.(color.RGBA)
	cr := (c2.R - c1.R)
	s := uint(cr * cr)
	cg := (c2.G - c1.G)
	s += uint(cg * cg)
	cb := (c2.B - c1.B)
	s += uint(cb * cb)
	return s
}

// Convert an image from truecolor to a N color palette
func Convert(img image.Image, pal color.Palette) image.Image {
	// For each pixel, go through each color in the palette and
	fmt.Println("CONVERT!")
	return img
}
