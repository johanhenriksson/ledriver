#include <stdbool.h>
#include <FastLED.h>
#include "common.h"
#include "controller.h"

static bool _initialized = false;
static int _count = 0;
static CRGB _leds[MAX_LEDS];
static CLEDController *_controller;

void led_init() {
    if (_initialized) {
        return;
    }

    _controller = &FastLED.addLeds<WS2811, DATA_PIN>(_leds, MAX_LEDS);
    _initialized = true;
}

void led_setup(int count) {
    if (_initialized) {
        _controller->setLeds(_leds, count);
        _controller->clearLeds(count);
    }
}

void led_set(int idx, CRGB color) {
    if (idx >= 0 && idx < MAX_LEDS) {
        _leds[idx] = color;
    }
}

void led_clear(CRGB color) {
    for(int i = 0; i < MAX_LEDS; i++) {
        _leds[i] = color;
    }
}

void led_show() {
    if (_initialized) {
        _controller->showLeds();
    }
}

void led_flash(CRGB color, int duration) {
    led_clear(color);
    led_show();
    delay(duration);
    led_clear(CRGB::Black);
    led_show();
}

