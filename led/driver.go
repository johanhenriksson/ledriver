package led

import (
    "fmt"
    "github.com/distributed/sers"
)

const (
    READ_BUFFER_SIZE = 64
    DEVICE_STRING = "/dev/cu.usbmodemfd%d"

    // messages
    MSG_SETUP       byte = 0x01
    MSG_FLIP_BUFFER byte = 0x02
    MSG_SET_PIXELS  byte = 0x10
)

// represents a serial connection to the microcontroller
type Driver struct {
    Serial  sers.SerialPort
    Rate    int
}

func NewDriver(id, rate int) *Driver {
    // open serial connection
    con, err := sers.Open(fmt.Sprintf(DEVICE_STRING, id))
    if err != nil {
        panic(fmt.Sprintf("Cannot open serial port: %s\n", err))
    }

    // set connection mode
    err = con.SetMode(rate, 8, sers.N, 2, sers.RTSCTS_HANDSHAKE)
    if err != nil {
        panic(fmt.Sprintf("Cannot set mode: %s\n", err))
    }

    fmt.Printf("Serial connection opened\n")

    d := &Driver {
        Serial: con,
        Rate: rate,
    }

    d.readInit()

    return d
}

func (d *Driver) Setup(id, width, height, array_width, array_height int) {
    m := []byte {
        MSG_SETUP,
        byte(id),
        byte(width),
        byte(height),
        byte(array_width),
        byte(array_height),
    }
    d.Serial.Write(m)
    d.readAck()
}

func (d *Driver) Show() {
    m := []byte { MSG_FLIP_BUFFER }
    d.Serial.Write(m)
    d.readAck()
}

// Write pixel data to the controller
func (d *Driver) SetPixels(start int, data []Color, auto_show bool) {
    const header_length = 4
    index := byte(start)
    count := byte(len(data))
    m := make([]byte, header_length + 3 * int(count) + 1)

    // header
    m[0] = MSG_SET_PIXELS
    if auto_show { m[1] = 1 }
    m[2] = index
    m[3] = count

    // pixel data payload
    i := header_length
    for _, color := range data {
        output := color.Output()
        m[i + 0] = output.R
        m[i + 1] = output.G
        m[i + 2] = output.B
        i += 3
    }

    // end
    m[i] = 0xFF

    d.Serial.Write(m)
    d.readAck()
}

// Shorthand for setting a single pixel
func (d *Driver) SetPixel(index int, color Color, auto_show bool) {
    d.SetPixels(index, []Color { color }, auto_show)
}

func (d *Driver) readByte() byte {
    buf := []byte { 0 }
    n, err := d.Serial.Read(buf)
    if err != nil {
        panic(fmt.Sprintf("Read error: %s\n", err))
    }
    if n != 1 {
        panic("Could not read byte")
    }
    return buf[0]
}

func (d *Driver) readLine() string {
    buf := make([]byte, READ_BUFFER_SIZE)
    buf[0] = d.readByte()

    i := 1
    for buf[i-1] != '\n' {
        buf[i] = d.readByte()
        i += 1
    }

    return string(buf[:i])
}

func (d *Driver) readAck() {
    fmt.Println("waiting for ack")
    b := d.readByte()
    if b != 0x01 {
        if b == 0xF0 {
            // seems to be a leftover init message
            fmt.Println("0xF0")
            fmt.Printf("%x\n",d.readByte())
            fmt.Print("Leftover data:", d.readLine())
            return
        }
        if b == 0xFF {
            // read error message
            panic(fmt.Sprintf("Driver error: %s", d.readLine()))
        }
        panic(fmt.Sprintf("Ack failed. Read %X", b))
    }
}

func (d *Driver) readInit() {
    // skip any trash
    var b, lb byte
    for !(b == 0xFA && lb == 0xF0) {
        lb = b
        b = d.readByte()
    }
    fmt.Print("<<", d.readLine())
}