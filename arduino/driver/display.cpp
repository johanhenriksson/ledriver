#include "display.h"
#include "state.h"

extern t_screen_state section;

int positionIndex(int x, int y) {
  if (y % 2 == 1) {
    x = section.width -x - 1;
  }
  return y * section.width + x;
}
