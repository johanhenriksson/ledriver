package led

import (
    "fmt"
    "github.com/tarm/serial"
)

const (
    READ_BUFFER_SIZE = 64
    SERIAL_RATE = 115200

    // messages
    MSG_START       byte = 0x02
    MSG_END         byte = 0x03
    MSG_HELLO       byte = 0xF0

    MSG_SETUP       byte = 0x01
    MSG_FLIP_BUFFER byte = 0x02
    MSG_SET_PIXELS  byte = 0x10
)

// represents a serial connection to the microcontroller
type Driver struct {
    Serial  *serial.Port
    Name    string
}

func NewDriver(device string) *Driver {
    conf := &serial.Config {
        Name: device,
        Baud: SERIAL_RATE,
    }

    // open serial connection
    con, err := serial.OpenPort(conf)
    if err != nil {
        panic(fmt.Sprintf("Cannot open serial port: %s\n", err))
    }

    fmt.Println("Opened", device)

    d := &Driver {
        Name: device,
        Serial: con,
    }

    d.readInit()

    fmt.Println("Device", device, "initialized")

    return d
}

func (d *Driver) Setup(id, width, height, array_width, array_height int) {
    m := []byte {
        MSG_START,
        MSG_SETUP,
        byte(id),
        byte(width),
        byte(height),
        byte(array_width),
        byte(array_height),
        MSG_END,
    }
    d.Serial.Write(m)
    d.readAck()
}

func (d *Driver) Show() {
    m := []byte { 
        MSG_START,
        MSG_FLIP_BUFFER,
        MSG_END,
    }
    d.Serial.Write(m)
    d.readAck()
}

// Write pixel data to the controller
func (d *Driver) SetPixels(start int, data []Color, auto_show bool) {
    const header_length = 5
    index := byte(start)
    count := byte(len(data))
    m := make([]byte, header_length + 3 * int(count) + 1)

    // header
    m[0] = MSG_START
    m[1] = MSG_SET_PIXELS
    if auto_show { m[2] = 1 }
    m[3] = index
    m[4] = count

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
    m[i] = MSG_END

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
    b := d.readByte()
    if b != 0x01 {
        if b == 0xF0 {
            // seems to be a leftover init message
            d.readByte()
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
    /*
    d.Serial.Write([]byte {
        MSG_START,
        MSG_HELLO,
        MSG_END,
    })
    fmt.Println("Wrote handshake")
    */

    // skip any trash
    var b, lb byte
    for !(b == 0xFA && lb == 0xF0) {
        lb = b
        b = d.readByte()
        fmt.Printf("Byte %x\n", b)
    }
}
