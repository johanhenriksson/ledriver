import serial
from typing import List

BAUD_RATE = 115200

MSG_START = 0x02
MSG_END = 0x03
MSG_ACK = 0x01
MSG_ERROR = 0xFF

MSG_SHOW = 0x02
MSG_PIXELS = 0x10

class DriverError(RuntimeError):
    pass

class Driver:
    def __init__(self, device: str):
        self.device = device
        self._serial = serial.Serial(device, BAUD_RATE)
        if self._readbyte() != MSG_START:
            raise DriverError('Expected MSG_START on init')

    def _readbyte(self) -> int:
        r = self._serial.read(1)
        return int(r[0])

    def _command(self, command: int, *data) -> None:
        self._serial.write(bytes([
            MSG_START,
            command,
            *data,
            MSG_END,
        ]))
        self._serial.flush()
        r = self._readbyte()
        if r == MSG_ACK:
            return
        elif r == MSG_ERROR:
            errlen = self._readbyte()
            errmsg = self._serial.read(errlen)
            raise DriverError(errmsg.decode('utf8'))
        else:
            raise DriverError('invalid response')

    def show(self) -> None:
        self._command(MSG_SHOW)

    def set(self, offset: int, colors: List[int]) -> None:
        if len(colors) % 3 != 0:
            raise ValueError('illegal number of color values')
        self._command(MSG_PIXELS, offset, len(colors) // 3, *colors)


if __name__ == '__main__':
    port = '/dev/tty.usbserial-22130'
    driver = Driver(port)
    driver.set(0, [
        255, 0, 0,
        0, 255, 0,
        0, 0, 255,
    ])
    driver.show()

