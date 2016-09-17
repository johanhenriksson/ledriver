package gfx

type Buffer []Color

func (b Buffer) Each(f func(*Color)) {
    for i := 0; i < len(b); i++ {
        f(&b[i])
    }
}

func (b Buffer) Add(color Color) {
    for i := 0; i < len(b); i++ {
        b[i] = b[i].Add(color)
    }
}

func (b Buffer) Multf(scalar float64) {
    for i := 0; i < len(b); i++ {
        b[i] = b[i].Multf(scalar)
    }
}

func (b Buffer) Mult(color Color) {
    for i := 0; i < len(b); i++ {
        b[i] =b[i].Mult(color)
    }
}

func (b Buffer) Clear(color Color) {
    for i := 0; i < len(b); i++ {
        b[i] = color
    }
}
