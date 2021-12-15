#ifndef COMMON_H
#define COMMON_H

#define MSG_START       0x02
#define MSG_END         0x03

// From client
#define MSG_SETUP       0x01
#define MSG_FLIP_BUFFER 0x02
#define MSG_SET_PIXELS  0x10
#define MSG_HELLO       0xF0

// To client
#define MSG_ACK         0x01
#define MSG_ERROR       0xFF

#define DATA_PIN        13
#define MAX_LEDS        128
#define SERIAL_RATE     115200
#define DRIVER_VERSION  "ledriver 1.0"

#endif
