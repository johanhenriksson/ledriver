#ifndef SERIAL_H
#define SERIAL_H

byte serial_read();
void serial_ack();
void serial_error(const char* format, ...);
void serial_init();
void serial_handshake();

#endif
