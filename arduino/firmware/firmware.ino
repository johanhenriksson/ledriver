#include <FastLED.h>

#include "common.h"
#include "controller.h"
#include "serial.h"

void handle_setup();
void handle_show();
void handle_pixels();

void setup() {
    led_init();

    // start by flashing all leds max green
    led_flash(CRGB::Red, 500);
    led_flash(CRGB::Green, 500);
    led_flash(CRGB::Blue, 500);

    // setup serial connection
    serial_init();
}

void loop() {  
    // read bytes until we find the start of a message
    while(serial_read() != MSG_START) { }

    // when we find the start of a message, the next byte will be the command
    byte cmd = serial_read();
    
    switch(cmd) {
    case MSG_SETUP:
        handle_setup();
        break;

    case MSG_SHOW: 
        handle_show();
        break;
    
    case MSG_PIXELS:
        handle_pixels();
        break;
    }

    // verify packet ending
    byte ending = serial_read();
    if (ending != MSG_END) {
        serial_error("invalid packet ending");
        return;
    }

    // acknowledge command
    serial_ack();
}

void handle_setup() {
    byte count = serial_read();
    led_setup(count);
}

void handle_show() {
    led_show();
}

void handle_pixels() {
    byte start = serial_read();
    byte count = serial_read();

    // cap at max leds. mostly in case the data is bad
    if (start >= MAX_LEDS) {
        return;
    }
    if (count >= MAX_LEDS) {
        return;
    }
    if (start + count > MAX_LEDS) {
        count = MAX_LEDS - start;
    }

    // write color data
    for (int i = 0; i < count; i++) {
        CRGB color;
        color.r = serial_read();
        color.g = serial_read();
        color.b = serial_read();
        led_set(start + i, color);
    }
}
