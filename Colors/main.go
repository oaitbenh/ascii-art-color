package Colors

import (
	"errors"
	"strconv"
	"strings"
)

type Color struct {
	R, G, B int
}

var Colors = map[string]Color{
	"reset":    {R: 255, G: 255, B: 255},
	"Black":    {R: 0, G: 0, B: 0},
	"White":    {R: 255, G: 255, B: 255},
	"Red":      {R: 255, G: 0, B: 0},
	"Green":    {R: 0, G: 255, B: 0},
	"Blue":     {R: 0, G: 0, B: 255},
	"Yellow":   {R: 255, G: 255, B: 0},
	"Magenta":  {R: 255, G: 0, B: 255},
	"Cyan":     {R: 0, G: 255, B: 255},
	"Orange":   {R: 255, G: 165, B: 0},
	"Purple":   {R: 128, G: 0, B: 128},
	"Brown":    {R: 165, G: 42, B: 42},
	"Lime":     {R: 191, G: 255, B: 0},
	"Pink":     {R: 255, G: 105, B: 180},
	"Teal":     {R: 0, G: 128, B: 128},
	"Lavender": {R: 230, G: 230, B: 250},
	"Beige":    {R: 245, G: 245, B: 220},
	"Maroon":   {R: 128, G: 0, B: 0},
	"Navy":     {R: 0, G: 0, B: 128},
	"Aqua":     {R: 127, G: 255, B: 212},
	"Olive":    {R: 128, G: 128, B: 0},
	"Gray":     {R: 128, G: 128, B: 128},
	"Silver":   {R: 192, G: 192, B: 192},
}

// func main() {
// 	Clr := "95c5e8"
// 	color := ""
// 	if strings.HasPrefix(Clr, "#") {
// 		color, _ = HexToRgb(Clr)
// 	}
// 	fmt.Printf(color + "Hello")
// }

func RgbColor(r, g, b int) string {
	R := strconv.Itoa(r)
	G := strconv.Itoa(g)
	B := strconv.Itoa(b)
	return "\033[38;2;" + R + ";" + G + ";" + B + "m"
}

func HexToRgb(Hex string) (string, error) {
	Hex = strings.TrimPrefix(Hex, "#")
	var r, g, b string
	switch len(Hex) {
	case 6:
		r = Hex[:2]
		g = Hex[2:4]
		b = Hex[4:]
	case 3:
		r = string(Hex[0])
		g = string(Hex[1])
		b = string(Hex[2])
	}
	R, err := strconv.ParseUint(r, 16, 8)
	if err != nil {
		return "", errors.New("This is not a hex Color")
	}
	G, err := strconv.ParseUint(g, 16, 8)
	if err != nil {
		return "", errors.New("This is not a hex Color")
	}
	B, err := strconv.ParseUint(b, 16, 8)
	if err != nil {
		return "", errors.New("This is not a hex Color")
	}
	return ("\033[38;2;" + strconv.Itoa(int(R)) + ";" + strconv.Itoa(int(G)) + ";" + strconv.Itoa(int(B)) + "m"), nil
}
