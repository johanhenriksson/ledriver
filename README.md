# ledriver

**Driver software for a LED screen project.**

Consists of two parts: the Arduino driver software and the PC control software.

The Arduino driver uses the FastLED library to control WS2811 LEDs over SPI, and implements a simple serial protocol that can recieve commands to display data on the screen. In order to achieve maximum performance, the Arduino does as little work as possible.

The control software implements a client for the serial protocol as well as functionality to handle square displays, screen buffering and gamma correction. More to come.

License: MIT
