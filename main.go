package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"

    "github.com/johanhenriksson/ledriver/led"
)

const (
    SERIAL_RATE = 250000
    MAX_FPS = 30
)

func main() {
    var id int
    fmt.Printf("Enter USB device id: ")
    //fmt.Scanf("%d", &id)
    fmt.Println()
    id = 1241

    driver := led.NewDriver(id, SERIAL_RATE)
    screen := led.NewScreen(driver, 2, 10, 10)
    frame_time := time.Second / MAX_FPS

    start := time.Now()
    frames := 500
    i := 0
    for  {
        frame_start := time.Now()
        screen.Clear(led.Color{0,0,0,1})

        chance := float32(0.80)
        f := float32(1.0)
        for rand.Float32() > (1.0 - chance * f) {
            f = (0.5 + 0.5 * (float32(math.Sin(float64(i) / 10.0)) + 1.0) / 2.0)
            intensity := rand.Float32() * f * 0.8
            color := led.FromHSL(float32((i + rand.Intn(100)) % frames) / float32(frames), 0.75, intensity)
            screen.Set(led.Point { byte(rand.Intn(10)), byte(rand.Intn(10)) }, color)
        }
        //p := byte(i % 100)
        //screen.Set(led.Point { p % 10, p / 10 }, led.Color { 0,1,0,1 })
        screen.Flip()
        frame_elapsed := time.Since(frame_start)

        time.Sleep(frame_time - frame_elapsed - time.Millisecond)
        i++
    }
    elapsed := time.Since(start)
    fmt.Println("1000 frames in", elapsed)
    fmt.Println("FPS:", float64(frames) / elapsed.Seconds())
}
