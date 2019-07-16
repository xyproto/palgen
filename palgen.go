package palgen

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

// SortablePalette is a slice of color.Color that can be sorted with sort.Sort, by euclidian distance for R, G and B
type SortablePalette []color.Color

// Length from RGB (0, 0, 0)
func colorLength(c color.Color) float64 {
	r := c.(color.RGBA)
	return math.Sqrt(float64(r.R*r.R + r.G*r.G + r.B*r.B)) // + r.A*r.A))
}

func (a SortablePalette) Len() int           { return len(a) }
func (a SortablePalette) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortablePalette) Less(i, j int) bool { return colorLength(a[i]) < colorLength(a[j]) }

// Median finds not the average but the median color
func Median(colors []color.Color) (color.Color, error) {
	if len(colors) == 0 {
		return nil, errors.New("can't find the median of an empty slice of colors")
	}
	if len(colors) == 1 {
		return colors[0], nil
	}

	// 1. Sort the colors
	sp := SortablePalette(colors)
	sort.Sort(sp)

	// 2. Select the center one, if odd
	if len(sp)%2 != 0 {
		centerPos := len(sp) / 2
		return sp[centerPos], nil
	}
	// 3. If the numbers are even, select the two center one and take the average of those
	centerPos1 := (len(sp) / 2) - 1
	centerPos2 := len(sp) / 2
	c1 := sp[centerPos1].(color.RGBA)
	c2 := sp[centerPos2].(color.RGBA)
	r := (c1.R + c2.R) / 2.0
	g := (c1.G + c2.G) / 2.0
	b := (c1.B + c2.B) / 2.0
	a := (c1.A + c2.A) / 2.0
	// return the new color
	return color.RGBA{r, g, b, a}, nil
}

// Generate can generate a palette with N colors, given an image
func Generate(img image.Image, N int) (color.Palette, error) {
	groups := make(map[int][]color.Color)
	already := make(map[color.Color]bool)

	// Pick out the colors from the image, per intensity level, and store them in the groups map
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y)
			gc := color.GrayModel.Convert(c).(color.Gray)
			rgba := color.RGBAModel.Convert(c).(color.RGBA)
			level := int(float64(gc.Y) / (255.0 / float64(N-1)))
			alreadyColor, ok := already[rgba]
			if !alreadyColor || !ok {
				groups[level] = append(groups[level], rgba)
				already[rgba] = true
			}
		}
	}

	//fmt.Println(groups)

	var pal color.Palette
	for _, colors := range groups {
		//fmt.Printf("Colors, group %d:\n", i)
		medianColor, err := Median(colors)
		if err != nil {
			return nil, err
		}
		//fmt.Println("\tmedian color", medianColor)
		pal = append(pal, medianColor)
	}

	return pal, nil
}

// GPL converts a given palette to the GIMP Palette Format (.gpl)
// The given name will be used as the palette name in the header
func GPL(pal color.Palette, name string) string {
	var sb strings.Builder
	// Prepare a header
	sb.WriteString("GIMP Palette\n")
	sb.WriteString("Name: ")
	sb.WriteString(name)
	sb.WriteString("\n")
	sb.WriteString("Columns: 4\n")
	sb.WriteString("# xyproto/palgen\n")
	// Output the colors
	for i, c := range pal {
		cn := c.(color.RGBA)
		sb.WriteString(fmt.Sprintf("%3d %3d %3d\t%d\n", cn.R, cn.G, cn.B, i))
	}
	// Return the generated string
	return sb.String()
}

// Save a palette to file in the GIMP Palette Format (.gpl)
// The given name will be used as the palette name in the header
func Save(pal color.Palette, filename, name string) error {
	return ioutil.WriteFile(filename, []byte(GPL(pal, name)), 0644)
}
