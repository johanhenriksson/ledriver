package led

type Screen struct {
    Index      int
    Width      int
    Height     int

    buffer     []Color
    driver     *Driver
}

func NewScreen(driver *Driver, index, width, height, array_width, array_height int) *Screen {
    s := &Screen {
        Index:  index,
        Width:  width,
        Height: height,

        driver: driver,
        buffer: make([]Color, width * height),
    }
    driver.Setup(index, width, height, array_width, array_height)
    s.Clear(Color{0,0,0,1})
    return s
}

// Clear screen buffer
func (s *Screen) Clear(color Color) {
    for i := 0; i < len(s.buffer); i++ {
        s.buffer[i] = color
    }
}

// Flip screen buffer
func (s *Screen) Flip() {
    s.driver.SetPixels(0, s.buffer, true)
}

func (s *Screen) Set(pos Vec, color Color) {
    s.buffer[s.index(pos)] = color
}

func (s *Screen) index(point Vec) int {
    x := int(point.X)
    if point.Y % 2 == 1 {
        x = s.Width - x - 1
    }
    return int(point.Y) * s.Width + x
}

