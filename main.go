package main

import (
    "fmt"
    "github.com/tarm/goserial"
)

const (
    SERIAL_RATE = 250000
)

func main() {
    c := &serial.Config {
        Name: fmt.Sprintf("/dev/usbmodem%d", id),
        Baud: SERIAL_RATE,
    }

    fmt.Println("ledriver")
}
