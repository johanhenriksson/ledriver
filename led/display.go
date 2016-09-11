package led
// screen arary

import (
    "fmt"
)

type DrawBuffer interface {
    Flip()
    Clear(color Color)
    Set(pos Vec, color Color)
}

type Display struct {
    ArrayWidth int
    ArrayHeight int
    ScreenWidth int
    ScreenHeight int
    Width int
    Height int
    Screens []*Screen
}

func NewDisplay(array_width, array_height int, screen_width, screen_height int) *Display {
    count := array_width * array_height
    d := &Display {
        ArrayWidth: array_width,
        ArrayHeight: array_height,
        ScreenWidth: screen_width,
        ScreenHeight: screen_height,
        Width: array_width * screen_width,
        Height: array_height * screen_height,
        Screens: make([]*Screen, count, count),
    }
    return d
}

func (d *Display) AttachScreen(index int, driver *Driver) {
    screen := NewScreen(driver, index, d.ScreenWidth, d.ScreenWidth, d.ArrayWidth, d.ArrayHeight)
    d.Screens[index] = screen;
    fmt.Println("Attach", driver.Name, "to index", index)
}

func (d *Display) Set(pos Vec, color Color) {
    // index?
    ix := int(pos.X) / d.ScreenWidth
    iy := int(pos.Y) / d.ScreenHeight
    sx := int(pos.X) % d.ScreenWidth
    sy := int(pos.Y) % d.ScreenHeight

    idx := iy * d.ArrayWidth + ix
    if d.Screens[idx] != nil {
        //panic("No such screen")
        d.Screens[idx].Set(Vec { byte(sx), byte(sy) }, color)
    }

}

func (d *Display) Flip() {
    for _, screen := range d.Screens {
        if screen != nil {
            screen.Flip()
        }
    }
}

func (d *Display) Clear(color Color) {
    for _, screen := range d.Screens {
        if screen != nil {
            screen.Clear(color)
        }
    }
}
