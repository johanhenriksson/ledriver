package gfx

import (
    "math"
    "math/rand"
)

type Effect interface {
    Draw(Buffer)
    Tick(int)
}

// An array of effects implements the Effect interface :)
type Effects []Effect

func (effects Effects) Draw(buffer Buffer) {
    for _, effect := range effects {
        effect.Draw(buffer)
    }
}

func (effects Effects) Tick(beat int) {
    for _, effect := range effects {
        effect.Tick(beat)
    }
}

// solid color

type SolidColor struct {
    Color       ColorPicker
}

func (sc SolidColor) Tick(beat int) { }

func (sc SolidColor) Draw(buffer Buffer) {
    buffer.Clear(sc.Color.Next())
}

// multiply color

type Multiply struct {
    Color       ColorPicker
}

func (mp Multiply) Tick(beat int) { }
func (mp Multiply) Draw(buffer Buffer) {
    buffer.Mult(mp.Color.Next())
}

// flash/strobe

type Flash struct {
    Color       ColorPicker
    Length      float64

    color       Color
    duration    float64
}

func (f *Flash) Tick(beat int) {
    if beat % 4 == 0 {
        f.duration = f.Length
        f.color = f.Color.Next()
    }
}

func (f *Flash) Draw(buffer Buffer) {
    if f.duration > 0 {
        progress := math.Max(f.duration / f.Length, 0.0)
        color := f.color.Multf(progress)
        buffer.Add(color)

        f.duration -= Time.Delta;
    }
}

// twinkle effect

type Twinkle struct {
    picker      ColorPicker
    stars       map[int]Color
    Falloff     float64
}

func NewTwinkle(picker ColorPicker) *Twinkle {
    return &Twinkle {
        Falloff: 0.995,
        picker: picker,
        stars: make(map[int]Color),
    }
}

func (t *Twinkle) Tick(beat int) {
    if beat % 1 == 0 {
        for i := 0; i <= rand.Intn(4); i++ {
            p := rand.Intn(16)
            t.stars[p] = t.picker.Next()
        }
    }
}

func (t *Twinkle) Draw(buffer Buffer) {
    for pos, color := range t.stars {
        c := color.Multf(0.9 + 0.1 * t.Falloff)
        buffer[pos] = buffer[pos].Add(c)
        t.stars[pos] = c
    }
}
