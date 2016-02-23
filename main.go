package main

import (
    "fmt"
    "time"

    "github.com/johanhenriksson/ledriver/led"
)

const (
    SERIAL_RATE = 250000
)

func main() {
    var id int
    fmt.Printf("Enter USB device id: ")
    fmt.Scanf("%d", &id)

    driver := led.NewDriver(id, SERIAL_RATE)
    driver.Fail()
    screen := led.NewScreen(driver, 2, 10, 10)

    start := time.Now()
    for i := 0; i < 1000; i++ {
        p := byte(i % 100)
        screen.Clear(led.Color{0,0,0,1})
        screen.Set(led.Point { p % 10, p / 10 }, led.Color{0,1,0,1})
        screen.Flip()
    }
    elapsed := time.Since(start)
    fmt.Println("1000 frames in", elapsed)
    fmt.Println("FPS:", 1000.0 / elapsed.Seconds())
}
