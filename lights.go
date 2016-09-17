package main

import (
    "time"
	"runtime"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/johanhenriksson/ledriver/gfx"
)

const SIZE = 16
const W = 4
const H = 4
const LIGHT_SIZE = 100

func main() {
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_EVERYTHING)

    scene := gfx.NewScene(SIZE)

    d := NewDisplay()
    d.Run(scene)

    sdl.Quit()

}

type Display struct {
	Ptr		*sdl.Window
}

func NewDisplay() *Display {
    window, err := sdl.CreateWindow("Lights", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 
        W * LIGHT_SIZE, H * LIGHT_SIZE, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
	return &Display {
        Ptr: window,
	}
}

func tickLoop(scene *gfx.Scene) {
    b := 0
    for {
        scene.Tick(b)
        time.Sleep(time.Second)
        b++
    }
}

func (wnd *Display) Run(scene *gfx.Scene) {
    surface, err := wnd.Ptr.GetSurface()
    if err != nil {
        panic(err)
    }

    go tickLoop(scene)

    running := true
	for running {
        var event sdl.Event
        for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                running = false
                break
            }
        }

        i := 0
        buffer := scene.Draw()
        for y := 0; y < H; y++ {
            for x := 0; x < W; x++ {
                pixel := buffer[i]
                rect := sdl.Rect{
                    int32(LIGHT_SIZE * x + 1), int32(LIGHT_SIZE * y + 1), // x, y
                    LIGHT_SIZE - 2, LIGHT_SIZE - 2, // w,h
                }

                clr := sdl.MapRGB(surface.Format,
                    uint8(255 * pixel.R()),
                    uint8(255 * pixel.G()),
                    uint8(255 * pixel.B()))

                surface.FillRect(&rect, clr)
                i++
            }
        }

        //scene.Buffer.Multf(0.98)

        wnd.Ptr.UpdateSurface()
	}
}
