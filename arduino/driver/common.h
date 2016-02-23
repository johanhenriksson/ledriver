#ifndef COMMON_H
#define COMMON_H

// From client
#define MSG_SETUP       0x01
#define MSG_FLIP_BUFFER 0x02
#define MSG_SET_PIXELS  0x10

// To client
#define MSG_ACK         0x01
#define MSG_ERROR       0xFF

#define DATA_PIN        2
#define MAX_LEDS        128
#define SERIAL_RATE 250000
#define DRIVER_VERSION  "ledriver prototype"

#define ROM_HAS_CONFIG     0
#define ROM_SECTION_ID     1
#define ROM_SECTION_WIDTH  2
#define ROM_SECTION_HEIGHT 3
#define ROM_ARRAY_WIDTH    4
#define ROM_ARRAY_HEIGHT   5

#endif
