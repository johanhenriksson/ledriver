#include <FastLED.h>
#include <EEPROM.h>

// From client
#define MSG_SETUP       0x01
#define MSG_FLIP_BUFFER 0x02
#define MSG_SET_PIXELS  0x10

// To client
#define MSG_ACK         0x01
#define MSG_ERROR       0xFF

#define DATA_PIN        2
#define MAX_LEDS        100
#define SERIAL_RATE     250000
#define DRIVER_VERSION  "ledriver prototype"

#define ROM_HAS_CONFIG     0
#define ROM_SECTION_ID     1
#define ROM_SECTION_WIDTH  2
#define ROM_SECTION_HEIGHT 3
#define ROM_ARRAY_WIDTH    4
#define ROM_ARRAY_HEIGHT   5

// globals
int  section_id = 0;
int  section_width = 0;
int  section_height = 0;
int  array_width = 3;
int  array_height = 2;
CRGB leds[MAX_LEDS];
char error_msg_buffer[128];

void setup() {
  // check for an on-chip stored configuration from a previous run
  if (load_config()) {
    identify();
  }

  // setup serial connection
  Serial.begin(250000);
  Serial.write(0xF0); // magic bytes
  Serial.write(0xFA);
  Serial.println(DRIVER_VERSION); // version string
  Serial.flush();
}

bool initialize(int section, int width, int height, int ar_width, int ar_height) {
  int count = width * height;
  if (count > MAX_LEDS) {
    error("Maximum number of leds: %d", MAX_LEDS);
    return false;
  }

  section_id = section;
  section_width = width;
  section_height = height;
  array_width = ar_width;
  array_height = ar_height;

  FastLED.clear();
  FastLED.addLeds<WS2811, DATA_PIN>(leds, count);
  return true;
}

// write current configuration to EEPROM
void save_config() {
  EEPROM.write(ROM_SECTION_ID,     section_id);
  EEPROM.write(ROM_SECTION_WIDTH,  section_width);
  EEPROM.write(ROM_SECTION_HEIGHT, section_height);
  EEPROM.write(ROM_ARRAY_WIDTH,    array_width);
  EEPROM.write(ROM_ARRAY_HEIGHT,   array_height);
  EEPROM.write(ROM_HAS_CONFIG,     1);
}

// load configuration from EEPROM
bool load_config() {
  int has_config = EEPROM.read(ROM_HAS_CONFIG);
  if (has_config) {
    // initialize WS2811 driver
    initialize(
      EEPROM.read(ROM_SECTION_ID),
      EEPROM.read(ROM_SECTION_WIDTH),
      EEPROM.read(ROM_SECTION_HEIGHT),
      EEPROM.read(ROM_ARRAY_WIDTH),
      EEPROM.read(ROM_ARRAY_HEIGHT)
    );
    return true;
  }
  return false;
}

void identify() {
  int i = 0;
  for(int y = 0; y < array_height; y++) {
    for(int x = 0; x < array_width; x++) {
      if (i == section_id) {
        leds[i] = CRGB::Green;
      } else {
        leds[i] = CRGB::Red;
      }
      i++;
    }
  }
  FastLED.show();
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
  Serial.write(MSG_ACK);
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

    Serial.write(MSG_ERROR);
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
      if (initialize(
        mustRead(), // section id
        mustRead(), // section width
        mustRead(), // section height
        mustRead(), // array width
        mustRead()  // array height
      )) {
        save_config();
        identify();
        ack();
      }
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
  }
}
