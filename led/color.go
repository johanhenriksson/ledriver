package led

import (
    "fmt"
    "os"
    "image"
)

type Color struct {
    R, G, B, A float32
}

// Gamma corrected output
type outputColor struct {
    R, G, B byte
}

func (c Color) Scale(f float32) Color {
    return Color {
        R: c.R * f,
        G: c.G * f,
        B: c.B * f,
        A: c.A * f,
    }
}
func (c Color) Output() outputColor {
    return outputColor {
        gamma_table[int(c.R * c.A * 255.99)],
        gamma_table[int(c.G * c.A * 255.99)],
        gamma_table[int(c.B * c.A * 255.99)],
    }
}

// stolen from https://github.com/bthomson/go-color
func FromHSL(h, s, l float32) Color {
    if s == 0 {
        // it's gray
        return Color{l, l, l, 1}
    }

    var v1, v2 float32
    if l < 0.5 {
        v2 = l * (1 + s)
    } else {
        v2 = (l + s) - (s * l)
    }

    v1 = 2*l - v2

    r := hueToRGB(v1, v2, h+(1.0/3.0))
    g := hueToRGB(v1, v2, h)
    b := hueToRGB(v1, v2, h-(1.0/3.0))

    return Color {r, g, b, 1}
}

// stolen from https://github.com/bthomson/go-color
func hueToRGB(v1, v2, h float32) float32 {
    if h < 0 {
        h += 1
    }
    if h > 1 {
        h -= 1
    }
    switch {
    case 6*h < 1:
        return (v1 + (v2-v1)*6*h)
    case 2*h < 1:
        return v2
    case 3*h < 2:
        return v1 + (v2-v1)*((2.0/3.0)-h)*6
    }
    return v1
}

func LoadImageColors(file string) {
    imgFile, err := os.Open(file)
    if err != nil {
        panic(fmt.Sprintf("Cannot open image %s: %s\n", file, err))
    }
    img, _, err := image.Decode(imgFile)
    if err != nil {
        panic(fmt.Sprintf("Cannot decode image %s: %s\n", file, err))
    }
    img.At(0,0)
}

var gamma_table []byte = []byte {
    0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
    0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  1,  1,  1,  1,
    1,  1,  1,  1,  1,  1,  1,  1,  1,  2,  2,  2,  2,  2,  2,  2,
    2,  3,  3,  3,  3,  3,  3,  3,  4,  4,  4,  4,  4,  5,  5,  5,
    5,  6,  6,  6,  6,  7,  7,  7,  7,  8,  8,  8,  9,  9,  9, 10,
    10, 10, 11, 11, 11, 12, 12, 13, 13, 13, 14, 14, 15, 15, 16, 16,
    17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 24, 24, 25,
    25, 26, 27, 27, 28, 29, 29, 30, 31, 32, 32, 33, 34, 35, 35, 36,
    37, 38, 39, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 50,
    51, 52, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 66, 67, 68,
    69, 70, 72, 73, 74, 75, 77, 78, 79, 81, 82, 83, 85, 86, 87, 89,
    90, 92, 93, 95, 96, 98, 99,101,102,104,105,107,109,110,112,114,
    115,117,119,120,122,124,126,127,129,131,133,135,137,138,140,142,
    144,146,148,150,152,154,156,158,160,162,164,167,169,171,173,175,
    177,180,182,184,186,189,191,193,196,198,200,203,205,208,210,213,
    215,218,220,223,225,228,231,233,236,239,241,244,247,249,252,255,
}
