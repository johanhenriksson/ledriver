#ifndef SERIAL_H
#define SERIAL_H

byte mustRead();
void ack();
void error(const char* format, ...);
void initSerial();
void handshake();

#endif
