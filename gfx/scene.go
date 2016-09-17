package gfx

import "time"

var Time = struct {
    Delta       float64
    Elapsed     float64
    lastFrame   time.Time
} { }

func init() {
    Time.lastFrame = time.Now()
}

type Scene struct {
    Effects Effects
    buffer  Buffer
}

func NewScene(size int) *Scene {
    g := LoadGradient("2")
    s := &Scene {
        buffer: make(Buffer, size),
        Effects: Effects {
            &SolidColor {
                Color: BLACK,
            },
            NewTwinkle(g),
            &Flash {
                Color: g,
                Length: 0.5,
            },
        },
    }
    return s
}

func (s Scene) Draw() Buffer {
    duration := time.Since(Time.lastFrame)
    Time.Delta = duration.Seconds()
    Time.Elapsed += Time.Delta
    Time.lastFrame = time.Now()

    s.Effects.Draw(s.buffer)
    return s.buffer
}

func (s Scene) Tick(beat int) {
    s.Effects.Tick(beat)
}
