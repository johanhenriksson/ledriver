#ifndef STATE_H
#define STATE_H

#include "Arduino.h"

// EEPROM addresses
#define ROM_HAS_STATE 0
#define ROM_STATE     1

typedef struct t_screen_state t_screen_state;

struct t_screen_state {
  byte id;
  byte width;
  byte height;
  byte array_width;
  byte array_height;
};

void rom_write(void* ptr, size_t count, int eeprom_dest);
void rom_read(int eeprom_src, size_t count, void* dest);

void state_init(t_screen_state* state);
bool state_set(t_screen_state *state, int id, int width, int height, int ar_width, int ar_height);
void state_save(t_screen_state* state);
bool state_load(t_screen_state* state);

#endif
