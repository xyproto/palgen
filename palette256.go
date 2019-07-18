package palgen

// An OK standard palette
var palette256 = [256][3]byte{
	[3]byte{0, 0, 0},
	[3]byte{0, 0, 102},
	[3]byte{0, 0, 204},
	[3]byte{0, 23, 51},
	[3]byte{0, 23, 153},
	[3]byte{0, 23, 255},
	[3]byte{0, 46, 0},
	[3]byte{0, 46, 102},
	[3]byte{0, 46, 204},
	[3]byte{0, 69, 51},
	[3]byte{0, 69, 153},
	[3]byte{0, 69, 255},
	[3]byte{0, 92, 0},
	[3]byte{0, 92, 102},
	[3]byte{0, 92, 204},
	[3]byte{0, 115, 51},
	[3]byte{0, 115, 153},
	[3]byte{0, 115, 255},
	[3]byte{0, 139, 0},
	[3]byte{0, 139, 102},
	[3]byte{0, 139, 204},
	[3]byte{0, 162, 51},
	[3]byte{0, 162, 153},
	[3]byte{0, 162, 255},
	[3]byte{0, 185, 0},
	[3]byte{0, 185, 102},
	[3]byte{0, 185, 204},
	[3]byte{0, 208, 51},
	[3]byte{0, 208, 153},
	[3]byte{0, 208, 255},
	[3]byte{0, 231, 0},
	[3]byte{0, 231, 102},
	[3]byte{0, 231, 204},
	[3]byte{0, 255, 51},
	[3]byte{0, 255, 153},
	[3]byte{0, 255, 255},
	[3]byte{42, 0, 51},
	[3]byte{42, 0, 153},
	[3]byte{42, 0, 255},
	[3]byte{42, 23, 0},
	[3]byte{42, 23, 102},
	[3]byte{42, 23, 204},
	[3]byte{42, 46, 51},
	[3]byte{42, 46, 153},
	[3]byte{42, 46, 255},
	[3]byte{42, 69, 0},
	[3]byte{42, 69, 102},
	[3]byte{42, 69, 204},
	[3]byte{42, 92, 51},
	[3]byte{42, 92, 153},
	[3]byte{42, 92, 255},
	[3]byte{42, 115, 0},
	[3]byte{42, 115, 102},
	[3]byte{42, 115, 204},
	[3]byte{42, 139, 51},
	[3]byte{42, 139, 153},
	[3]byte{42, 139, 255},
	[3]byte{42, 162, 0},
	[3]byte{42, 162, 102},
	[3]byte{42, 162, 204},
	[3]byte{42, 185, 51},
	[3]byte{42, 185, 153},
	[3]byte{42, 185, 255},
	[3]byte{42, 208, 0},
	[3]byte{42, 208, 102},
	[3]byte{42, 208, 204},
	[3]byte{42, 231, 51},
	[3]byte{42, 231, 153},
	[3]byte{42, 231, 255},
	[3]byte{42, 255, 0},
	[3]byte{42, 255, 102},
	[3]byte{42, 255, 204},
	[3]byte{85, 0, 0},
	[3]byte{85, 0, 102},
	[3]byte{85, 0, 204},
	[3]byte{85, 23, 51},
	[3]byte{85, 23, 153},
	[3]byte{85, 23, 255},
	[3]byte{85, 46, 0},
	[3]byte{85, 46, 102},
	[3]byte{85, 46, 204},
	[3]byte{85, 69, 51},
	[3]byte{85, 69, 153},
	[3]byte{85, 69, 255},
	[3]byte{85, 92, 0},
	[3]byte{85, 92, 102},
	[3]byte{85, 92, 204},
	[3]byte{85, 115, 51},
	[3]byte{85, 115, 153},
	[3]byte{85, 115, 255},
	[3]byte{85, 139, 0},
	[3]byte{85, 139, 102},
	[3]byte{85, 139, 204},
	[3]byte{85, 162, 51},
	[3]byte{85, 162, 153},
	[3]byte{85, 162, 255},
	[3]byte{85, 185, 0},
	[3]byte{85, 185, 102},
	[3]byte{85, 185, 204},
	[3]byte{85, 208, 51},
	[3]byte{85, 208, 153},
	[3]byte{85, 208, 255},
	[3]byte{85, 231, 0},
	[3]byte{85, 231, 102},
	[3]byte{85, 231, 204},
	[3]byte{85, 255, 51},
	[3]byte{85, 255, 153},
	[3]byte{85, 255, 255},
	[3]byte{127, 0, 51},
	[3]byte{127, 0, 153},
	[3]byte{127, 0, 255},
	[3]byte{127, 23, 0},
	[3]byte{127, 23, 102},
	[3]byte{127, 23, 204},
	[3]byte{127, 46, 51},
	[3]byte{127, 46, 153},
	[3]byte{127, 46, 255},
	[3]byte{127, 69, 0},
	[3]byte{127, 69, 102},
	[3]byte{127, 69, 204},
	[3]byte{127, 92, 51},
	[3]byte{127, 92, 153},
	[3]byte{127, 92, 255},
	[3]byte{127, 115, 0},
	[3]byte{127, 115, 102},
	[3]byte{127, 115, 204},
	[3]byte{127, 139, 51},
	[3]byte{127, 139, 153},
	[3]byte{127, 139, 255},
	[3]byte{127, 162, 0},
	[3]byte{127, 162, 102},
	[3]byte{127, 162, 204},
	[3]byte{127, 185, 51},
	[3]byte{127, 185, 153},
	[3]byte{127, 185, 255},
	[3]byte{127, 208, 0},
	[3]byte{127, 208, 102},
	[3]byte{127, 208, 204},
	[3]byte{127, 231, 51},
	[3]byte{127, 231, 153},
	[3]byte{127, 231, 255},
	[3]byte{127, 255, 0},
	[3]byte{127, 255, 102},
	[3]byte{127, 255, 204},
	[3]byte{170, 0, 0},
	[3]byte{170, 0, 102},
	[3]byte{170, 0, 204},
	[3]byte{170, 23, 51},
	[3]byte{170, 23, 153},
	[3]byte{170, 23, 255},
	[3]byte{170, 46, 0},
	[3]byte{170, 46, 102},
	[3]byte{170, 46, 204},
	[3]byte{170, 69, 51},
	[3]byte{170, 69, 153},
	[3]byte{170, 69, 255},
	[3]byte{170, 92, 0},
	[3]byte{170, 92, 102},
	[3]byte{170, 92, 204},
	[3]byte{170, 115, 51},
	[3]byte{170, 115, 153},
	[3]byte{170, 115, 255},
	[3]byte{170, 139, 0},
	[3]byte{170, 139, 102},
	[3]byte{170, 139, 204},
	[3]byte{170, 162, 51},
	[3]byte{170, 162, 153},
	[3]byte{170, 162, 255},
	[3]byte{170, 185, 0},
	[3]byte{170, 185, 102},
	[3]byte{170, 185, 204},
	[3]byte{170, 208, 51},
	[3]byte{170, 208, 153},
	[3]byte{170, 208, 255},
	[3]byte{170, 231, 0},
	[3]byte{170, 231, 102},
	[3]byte{170, 231, 204},
	[3]byte{170, 255, 51},
	[3]byte{170, 255, 153},
	[3]byte{170, 255, 255},
	[3]byte{212, 0, 51},
	[3]byte{212, 0, 153},
	[3]byte{212, 0, 255},
	[3]byte{212, 23, 0},
	[3]byte{212, 23, 102},
	[3]byte{212, 23, 204},
	[3]byte{212, 46, 51},
	[3]byte{212, 46, 153},
	[3]byte{212, 46, 255},
	[3]byte{212, 69, 0},
	[3]byte{212, 69, 102},
	[3]byte{212, 69, 204},
	[3]byte{212, 92, 51},
	[3]byte{212, 92, 153},
	[3]byte{212, 92, 255},
	[3]byte{212, 115, 0},
	[3]byte{212, 115, 102},
	[3]byte{212, 115, 204},
	[3]byte{212, 139, 51},
	[3]byte{212, 139, 153},
	[3]byte{212, 139, 255},
	[3]byte{212, 162, 0},
	[3]byte{212, 162, 102},
	[3]byte{212, 162, 204},
	[3]byte{212, 185, 51},
	[3]byte{212, 185, 153},
	[3]byte{212, 185, 255},
	[3]byte{212, 208, 0},
	[3]byte{212, 208, 102},
	[3]byte{212, 208, 204},
	[3]byte{212, 231, 51},
	[3]byte{212, 231, 153},
	[3]byte{212, 231, 255},
	[3]byte{212, 255, 0},
	[3]byte{212, 255, 102},
	[3]byte{212, 255, 204},
	[3]byte{255, 0, 0},
	[3]byte{255, 0, 102},
	[3]byte{255, 0, 204},
	[3]byte{255, 23, 51},
	[3]byte{255, 23, 153},
	[3]byte{255, 23, 255},
	[3]byte{255, 46, 0},
	[3]byte{255, 46, 102},
	[3]byte{255, 46, 204},
	[3]byte{255, 69, 51},
	[3]byte{255, 69, 153},
	[3]byte{255, 69, 255},
	[3]byte{255, 92, 0},
	[3]byte{255, 92, 102},
	[3]byte{255, 92, 204},
	[3]byte{255, 115, 51},
	[3]byte{255, 115, 153},
	[3]byte{255, 115, 255},
	[3]byte{255, 139, 0},
	[3]byte{255, 139, 102},
	[3]byte{255, 139, 204},
	[3]byte{255, 162, 51},
	[3]byte{255, 162, 153},
	[3]byte{255, 162, 255},
	[3]byte{255, 185, 0},
	[3]byte{255, 185, 102},
	[3]byte{255, 185, 204},
	[3]byte{255, 208, 51},
	[3]byte{255, 208, 153},
	[3]byte{255, 208, 255},
	[3]byte{255, 231, 0},
	[3]byte{255, 231, 102},
	[3]byte{255, 231, 204},
	[3]byte{255, 255, 51},
	[3]byte{255, 255, 153},
	[3]byte{255, 255, 255},
	[3]byte{204, 204, 204},
	[3]byte{153, 153, 153},
	[3]byte{102, 102, 102},
	[3]byte{51, 51, 51},
}