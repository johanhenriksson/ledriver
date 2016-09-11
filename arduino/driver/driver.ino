#include <FastLED.h>
#include <EEPROM.h>

#include "common.h"
#include "serial.h"
#include "state.h"
#include "display.h"

void clear();
void indentify();
void flip_buffer();
void setupLeds(int count);

// globals
CRGB leds[MAX_LEDS];
t_screen_state section;

void setup() {
  // check for an on-chip stored configuration from a previous run
  state_init(&section);
  if (state_load(&section)) {
    setupLeds(section.width * section.height);
    identify();
  }

  for(int i = 0; i < MAX_LEDS; i++)
    leds[i] = CRGB::Black;
  flip_buffer();
  
  // setup serial connection
  initSerial();
}

void loop() {  
  // read command id
  byte cmd;
  do {
    cmd = mustRead();
  } while(cmd != MSG_START);
  // when we find the start of a message, the next byte will be the command
  cmd = mustRead();
    
  byte start, count, auto_show, ending;
  switch(cmd) 
  {
    case MSG_HELLO: {
      // allows reconnection
      handshake();
      byte end = mustRead();
      if (end != MSG_END) {
        error("Invalid flip buffer ending");
      }
      break;
    }
      
    case MSG_SETUP: {
      byte id = mustRead(), // section id
           w  = mustRead(), // section width
           h  = mustRead(), // section height
           aw = mustRead(), // array width
           ah = mustRead(); // array height
           
      if (state_set(&section, id, w, h, aw, ah)) {
        state_save(&section);
        
        setupLeds(section.width * section.height);
        identify();

        byte end = mustRead();
        if (end != MSG_END) {
          error("Invalid flip buffer ending");
        }
        
        ack();
      }
      break;
    }

    // flip buffers
    case MSG_FLIP_BUFFER: {
      flip_buffer();
      byte end = mustRead();
      if (end != MSG_END) {
        error("Invalid flip buffer ending");
      }
      
      ack();
      break;
    }
    
    case MSG_SET_PIXELS: {
      auto_show = mustRead() > 0;
      start = mustRead();
      count = mustRead();

      // cap at max leds. mostly in case the data is bad
      if (start >= MAX_LEDS) {
        break;
      }
      if (start + count > MAX_LEDS) {
        count = MAX_LEDS - start;
      }

      // write color data
      for (int i = 0; i < count; i++) {
        leds[start + i].r = mustRead();
        leds[start + i].g = mustRead();
        leds[start + i].b = mustRead();
      }

      // check ending before flipping the buffer. if the data
      // is corrupt then skip the frame if possible
      ending = mustRead();
      if (ending == MSG_END) {
        // auto flip buffer?
        if (auto_show)
          flip_buffer();
      } else {
        error("Invalid SetPixel ending");
      }
      
      ack();
      break;
    }
  }
}

void identify() {
  clear();
  
  int i = 0;
  for(int y = 0; y < section.array_height; y++) {
    for(int x = 0; x < section.array_width; x++) {
      int p = positionIndex(x,y);
      if (i == section.id) {
        leds[p] = CRGB::Green;
      } else {
        leds[p] = CRGB::Red;
      }
      i++;
    }
  }
  flip_buffer();
}

void flip_buffer() {
  FastLED.show();
}

void setupLeds(int count) {
  FastLED.clear();
  FastLED.addLeds<WS2811, DATA_PIN>(leds, count);
}

void clear() {
  for(int i = 0; i < MAX_LEDS; i++)
    leds[i] = CRGB::Black;
}

