#include "Arduino.h"
#include "serial.h"
#include "common.h"

#define ERROR_BUFFER_SIZE 128

char error_msg_buffer[ERROR_BUFFER_SIZE];

// initializes the serial connection
void serial_init() {
  Serial.begin(SERIAL_RATE);
  Serial.write(MSG_START);
}

// blocks until a byte is available on the serial connection, then returns it
byte serial_read() {
  // wait for data
  while(Serial.available() <= 0) { }
    
  // read byte
  return Serial.read();
}

// writes an ACK message to the serial port
void serial_ack() {
  Serial.write(MSG_ACK);
  Serial.flush();
}

// write an error to serial. works like printf
void serial_error(const char* format, ...) {
    va_list argptr;
    va_start(argptr, format);
    vsprintf(error_msg_buffer, format, argptr);
    va_end(argptr);

    Serial.write(MSG_ERROR);
    Serial.write(strlen(error_msg_buffer));
    Serial.print(error_msg_buffer);
    Serial.flush();
}
