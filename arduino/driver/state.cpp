#include <EEPROM.h>
#include <FastLED.h>
#include "Arduino.h"
#include "state.h"
#include "common.h"
#include "serial.h"

// set defaults
void state_init(t_screen_state* state) {
  state->id           = 0;
  state->width        = 10;
  state->height       = 10;
  state->array_width  = 1;
  state->array_height = 1;
}

bool state_set(t_screen_state *state, int id, int width, int height, int ar_width, int ar_height) {
  int count = width * height;
  if (count > MAX_LEDS) {
    error("Maximum number of leds: %d", MAX_LEDS);
    return false;
  }

  state->id     = id;
  state->width  = width;
  state->height = height;
  state->array_width  = ar_width;
  state->array_height = ar_height;

  return true;
}

// save to eeprom
void state_save(t_screen_state* state) {
  rom_write(state, sizeof(t_screen_state), ROM_STATE);
  EEPROM.write(ROM_HAS_STATE, 1);
}

// attempt to load from eeprom
bool state_load(t_screen_state* state) {
  return false;
  int has_state = EEPROM.read(ROM_HAS_STATE);
  if (has_state) {
    rom_read(ROM_STATE, sizeof(t_screen_state), state);
    return true;
  }
  return false;
}

// we can use this to write whatever we want from eeprom
void rom_write(void* ptr, size_t count, int eeprom_dest) {
  for(size_t i = 0; i < count; i++) {
    byte v = *((byte*)((size_t)ptr + i));
    EEPROM.write(eeprom_dest + i, v);
  }
}

// we can use this to read whatever we want from eeprom
void rom_read(int eeprom_src, size_t count, void* dest) {
  for(size_t i = 0; i < count; i++) {
    byte* p = (byte*)((size_t)dest + i);
    *p = EEPROM.read(eeprom_src + i);
  }
}
