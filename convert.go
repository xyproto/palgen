package palgen

import (
	"errors"
	"image"
	"image/color"
)

// TODO: Create a custom Paletted type, based on image/Paletted, that can take > 256 colors

// ConvertCustom can convert an image from truecolor to a <=256 color paletted image
func ConvertCustom(m image.Image, pal color.Palette) (image.Image, error) {
	if len(pal) > 256 {
		return nil, errors.New("can convert to a maximum of 256 colors")
	}
	palImg := image.NewPaletted(m.Bounds(), pal)
	// For each pixel, go through each color in the palette and pick out the closest one.
	for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
		for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
			sourceColor := m.At(x, y)
			colorIndex := uint8(pal.Index(sourceColor))
			palImg.SetColorIndex(x, y, colorIndex)
		}
	}
	return palImg, nil
}

// Convert an image from truecolor to a 256 color paletted image, with a custom palette
func Convert(m image.Image) (image.Image, error) {
	customPalette, err := Generate(m, 256)
	if err != nil {
		return nil, err
	}
	// This should never happen
	if len(customPalette) > 256 {
		return nil, errors.New("the generated palette has too many colors")
	}
	// Return a new Paletted image
	return ConvertCustom(m, customPalette)
}

// ConvertGeneral can convert an image from truecolor to a 256 color paletted image, with a general palette
func ConvertGeneral(m image.Image) (image.Image, error) {
	return ConvertCustom(m, GeneralPalette())
}
