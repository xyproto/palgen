package palgen

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"strings"
)

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
