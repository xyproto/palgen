package palgen

import (
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
	retimg := image.NewPaletted(img.Bounds(), pal)
	// For each pixel, go through each color in the palette and pick out the closest one.
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			imageColor := img.At(x, y)
			//imageColor := color.RGBAModel.Convert(c).(color.RGBA)
			newColor := imageColor
			minD := uint(9999)
			for _, paletteColor := range pal {
				d := quickDistance(imageColor, paletteColor)
				if d == 0 {
					// Break out of the loop if the distance is 0
					newColor = paletteColor
					break
				} else if d < minD {
					minD = d
					newColor = paletteColor
				}
			}
			retimg.Set(x, y, newColor)
		}
	}
	return retimg
}
