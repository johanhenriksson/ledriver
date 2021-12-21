#ifndef CONTROLLER_H
#define CONTROLLER_H 
#include <pixeltypes.h>

void led_init();
void led_setup(int count);
void led_set(int idx, CRGB color);
void led_clear(CRGB color);
void led_show();

void led_flash(CRGB color, int duration);

#endif
