#include <FastLED.h>
#include <EEPROM.h>

#define MSG_SETUP       0x01
#define MSG_FLIP_BUFFER 0x02
#define MSG_SET_PIXELS  0x10

#define DATA_PIN 2
#define MAX_LEDS 100
#define DRIVER_VERSION "ledriver prototype"

const int rom_led_count = 0;
const int rom_section_id = 1;

// globals
int  led_count = 0;
int  section_id = 0;
CRGB leds[MAX_LEDS];
char error_msg_buffer[128];

void setup() {
  // check for an on-chip stored configuration
  led_count = EEPROM.read(rom_led_count);
  if (led_count > 0) {
    section_id = EEPROM.read(rom_section_id);

    // initialize WS2811 driver
    FastLED.addLeds<WS2811, DATA_PIN>(leds, led_count);

    // show section id
    for(int i = 0; i < section_id; i++)
      leds[i] = CRGB::Red;
    FastLED.show();
  }

  // initialize serial connection
  Serial.begin(250000);
  Serial.println(DRIVER_VERSION);
}

byte mustRead() {
  // wait for data
  while(Serial.available() <= 0)
    ;
  // read byte
  return Serial.read();
}

// write ack message
void ack() {
  Serial.write(0x01);
  Serial.flush();
}

void flip_buffer() {
  FastLED.show();
}

void error(const char* format, ...) {
    va_list argptr;
    va_start(argptr, format);
    vsprintf(error_msg_buffer, format, argptr);
    va_end(argptr);

    Serial.write(0xFF);
    Serial.println(error_msg_buffer);
    Serial.flush();
}

void loop() {
  // read command id
  byte cmd = mustRead();
    
  byte start, count, auto_show, ending;
  switch(cmd) 
  {
    case MSG_SETUP: {
      // read new configuration and write it to EEPROM
      section_id = mustRead();
      led_count = mustRead();

      if (led_count > MAX_LEDS) {
        error("Maximum number of leds: %d", MAX_LEDS);
        return;  
      }
      
      EEPROM.write(rom_led_count, led_count);
      EEPROM.write(rom_section_id, section_id);

      // initialize WS2811 driver
      FastLED.clear();
      FastLED.addLeds<WS2811, DATA_PIN>(leds, led_count);
      ack();
      break;
    }

    // flip buffers
    case MSG_FLIP_BUFFER: {
      flip_buffer();
      ack();
      break;
    }
    
    case MSG_SET_PIXELS: {
      auto_show = mustRead() > 0;
      start = mustRead();
      count = mustRead();
      for (int i = 0; i < count; i++) {
        leds[start + i].r = mustRead();
        leds[start + i].g = mustRead();
        leds[start + i].b = mustRead();
      }
      ending = mustRead();

      // auto flip buffer?
      if (auto_show)
        flip_buffer();
      
      ack();
      break;
    }
    
    default: {
      error("Unknown command");      
      break;
    }
  }
}
