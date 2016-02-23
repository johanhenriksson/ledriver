#include "Arduino.h"
#include "serial.h"
#include "common.h"

#define ERROR_BUFFER_SIZE 128

char error_msg_buffer[ERROR_BUFFER_SIZE];

void initSerial() {
  Serial.begin(SERIAL_RATE);
  Serial.write(0xF0); // magic bytes
  Serial.write(0xFA);
  Serial.println(DRIVER_VERSION); // version string
  Serial.flush();
}

// blocks until a byte is available on the serial connection, then returns it
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

// write an error to serial. works like printf
void error(const char* format, ...) {
    va_list argptr;
    va_start(argptr, format);
    vsprintf(error_msg_buffer, format, argptr);
    va_end(argptr);

    Serial.write(MSG_ERROR);
    Serial.println(error_msg_buffer);
    Serial.flush();
}
