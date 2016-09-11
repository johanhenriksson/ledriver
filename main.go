package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
    //"os"
    //"bufio"

    "github.com/johanhenriksson/ledriver/led"
)

const (
    MAX_FPS = 60
)

func main() {
    /*
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter USB device 1 id: ")
    dev1, _ := reader.ReadString('\n')
    fmt.Printf("Enter USB device 2 id: ")
    dev2, _ := reader.ReadString('\n')
    */

    driver := led.NewDriver("/dev/tty.usbmodem1421")
    //driver2 := led.NewDriver("/dev/ttyUSB4")

    count := 128
    driver.Setup(0, count, 1, 1, 1)

    fmt.Println("Display ready")

    black := make([]led.Color, count)
    white := make([]led.Color, count)

    for i := 0; i < count; i++ {
        black[i] = led.Color{1,0,0,1}
        white[i] = led.Color{0,0,1,1}
    }

    for {
        driver.SetPixels(0, white, true)
        time.Sleep(time.Second)
        driver.SetPixels(0, black, true)
        time.Sleep(time.Second)
    }

    display := led.NewDisplay(1, 1, 11, 1)
    display.AttachScreen(0, driver)
    //display.AttachScreen(1, driver2);

    frame_time := time.Second / MAX_FPS

    effects := led.NewEffects()

    timescale := float32(30)
    start := time.Now()
    frames := 1500
    i := 0
    frame_start := time.Now()
    chance := float32(0.75)

    for {
        display.Clear(led.Color{1,1,1,1})
        display.Flip()
        time.Sleep(50 * time.Millisecond)
        display.Clear(led.Color{0,0,0,1})
        display.Flip()
        time.Sleep(50 * time.Millisecond)
    }

    for {
        c1 := led.FromHSL(float32(i%frames)/float32(frames), 0.5 + 0.5 * rand.Float32(), 0.75)
        display.Set(led.Vec { 0,0 }, c1)
        display.Set(led.Vec { 1,0 }, c1)
        display.Set(led.Vec { 2,0 }, c1)

        c2 := led.FromHSL(-1.0 * float32(i%frames)/float32(frames), 0.5 + 0.5 * rand.Float32(), 0.75)

        display.Set(led.Vec { 3,0 }, c2)
        display.Set(led.Vec { 4,0 }, c2)
        display.Set(led.Vec { 5,0 }, c2)

        display.Flip()
        time.Sleep(100 * time.Millisecond)
        i++
    }

    for {
        f := float32(1.0)
        for rand.Float32() > (1.0 - chance * f) {

            f = (0.5 + 0.5 * (float32(math.Sin(float64(i) / 10.0)) + 1.0) / 2.0)
            intensity := float32(math.Sin(float64(i) * 0.3)) //rand.Float32() * f * 0.8
            //color := led.FromHSL(float32((i + rand.Intn(60)) % frames) / float32(frames), 0.75, intensity)
            color := led.FromHSL(float32(i%frames)/float32(frames), intensity, 0.5 + 0.42 * rand.Float32())
            //color := led.Color { 0, 0, 0.2 + 0.8 * rand.Float32(), 1}
            pos := led.Vec { byte(rand.Intn(display.Width)), byte(rand.Intn(display.Height)) }

            effects.Add(led.NewSpark(pos, color, intensity, 2, 0.0 + rand.Float32() * 0.1))

            i++
        }

        last_frame := time.Since(frame_start)
        frame_start = time.Now()
        effects.Update(float32(last_frame.Seconds()) * timescale)
        effects.Draw(display)
        frame_elapsed := time.Since(frame_start) 

        time.Sleep(frame_time - frame_elapsed - time.Millisecond)
        fmt.Println("Effects:", effects.Count())
    }

    elapsed := time.Since(start)
    fmt.Println("1000 frames in", elapsed)
    fmt.Println("FPS:", float64(frames) / elapsed.Seconds())
}
