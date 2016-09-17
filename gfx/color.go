package gfx

import (
    "os"
    "fmt"
    "time"
    "math"
    "math/rand"

    "image"
    _ "image/png"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Color  [3]float64

func (c Color) R() float64 { return c[0] }
func (c Color) G() float64 { return c[1] }
func (c Color) B() float64 { return c[2] }

func RandomColor() Color {
    return Color {
        rand.Float64(),
        rand.Float64(),
        rand.Float64(),
    }
}

var RED   = Color { 1.0, 0.0, 0.0 }
var GREEN = Color { 0.0, 1.0, 0.0 }
var BLUE  = Color { 0.0, 0.0, 1.0 }
var GRAY  = Color { 0.5, 0.5, 0.5 }
var WHITE = Color { 1.0, 1.0, 1.0 }
var BLACK = Color { 0.0, 0.0, 0.0 }

func (c Color) Multf(scalar float64) Color {
    return Color {
        math.Min(c[0] * scalar, 1.0),
        math.Min(c[1] * scalar, 1.0),
        math.Min(c[2] * scalar, 1.0),
    }
}

func (c Color) Mult(other Color) Color {
    return Color {
        c[0] * other[0],
        c[1] * other[1],
        c[2] * other[2],
    }
}

func (c Color) Add(other Color) Color {
    return Color {
        math.Min(c[0] + other[0], 1.0),
        math.Min(c[1] + other[1], 1.0),
        math.Min(c[2] + other[2], 1.0),
    }
}

type ColorPicker interface {
    Next() Color
}

// we can use a normal color as a color picker :)
func (c Color) Next() Color {
    return c
}

type GradientPicker struct {
    img     image.Image
    min     int
    max     int
}

func (gp *GradientPicker) Next() Color {
    x := rand.Intn(gp.max)
    r, g, b, _ := gp.img.At(x, 0).RGBA()
    c := Color {
        float64(r) / 65535.0,
        float64(g) / 65535.0,
        float64(b) / 65535.0,
    }
    return c
}

func LoadGradient(name string) *GradientPicker {
	reader, err := os.Open(fmt.Sprintf("assets/gradient/%s.png", name))
    if err != nil { panic(err) }
	defer reader.Close()

    img, _, err := image.Decode(reader)
    if err != nil { panic(err) }

	bounds := img.Bounds()
    gp := &GradientPicker {
        img: img,
        min: bounds.Min.X,
        max: bounds.Max.X,
    }
    return gp
}
