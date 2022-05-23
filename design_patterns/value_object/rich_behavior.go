package value_object

import (
	"fmt"
	"math"
)

// VO delivers many different behaviors. Its main purpose is to provide a reach interface
func (c Color) ToBrighter() Color {
	return Color{
		Red:   math.Min(255, c.Red+10),
		Green: math.Min(255, c.Green+10),
		Blue:  math.Min(255, c.Blue+10),
	}
}

func (c Color) ToDarker() Color {
	return Color{
		Red:   math.Max(0, c.Red-10),
		Green: math.Max(0, c.Green-10),
		Blue:  math.Max(0, c.Blue-10),
	}
}

func (c Color) Combine(other Color) Color {
	return Color{
		Red:   math.Min(255, c.Red+other.Red),
		Green: math.Min(255, c.Green+other.Green),
		Blue:  math.Min(255, c.Blue+other.Blue),
	}
}

func (c Color) IsRed() bool {
	return c.Red == 255 && c.Green == 0 && c.Blue == 0
}

func (c Color) IsYellow() bool {
	return c.Red == 255 && c.Green == 255 && c.Blue == 0
}

func (c Color) IsMagenta() bool {
	return c.Red == 255 && c.Green == 0 && c.Blue == 255
}

func (c Color) ToCSS() string {
	return fmt.Sprintf(`rgb(%d, %d, %d)`, c.Red, c.Green, c.Blue)
}
