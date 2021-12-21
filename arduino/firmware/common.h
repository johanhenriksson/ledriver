#ifndef COMMON_H
#define COMMON_H

#define MSG_START       0x02
#define MSG_END         0x03

// From client
#define MSG_SETUP       0x01
#define MSG_SHOW        0x02
#define MSG_PIXELS      0x10

// To client
#define MSG_ACK         0x01
#define MSG_ERROR       0xFF

#define DATA_PIN        13
#define MAX_LEDS        128
#define SERIAL_RATE     115200
#define DRIVER_VERSION  "1.0"

#endif
