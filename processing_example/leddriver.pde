import processing.serial.*;

public class LedDriver
{
  // serial rate (in baud)
  protected final int  SERIAL_RATE = 115200;

  // command ids
  protected static final byte MSG_SETUP = 0x01;
  protected static final byte MSG_FLIP  = 0x02;
  protected static final byte MSG_SET   = 0x10;
  
  // message start/end magic bytes
  protected static final byte MSG_START = 0x02;
  protected static final byte MSG_END   = 0x03;

  private PApplet app;
  private Serial port;
  private int lights;
  
  public LedDriver(PApplet app, String portName) {
    this.app = app;
    this.lights = 0;
    this.open(portName);
  }
  
  public int Length() {
    return this.lights;
  }
  
  protected void open(String portName) {
    port = new Serial(app, portName, SERIAL_RATE);
    readInit();
    println("Serial connection ready. Speed: " + SERIAL_RATE + " baud");
  }

  public void Close() {
      if (port != null) {
          port.stop();
          port = null;
          println("Serial connection closed");
      }
  }
  
  /*
   * methods for reading data
   */
  
  // block until a byte is available, then read it
  protected int readByte() {
    if (port == null)
        throw new RuntimeException("Serial device not connected");

    while(port.available() < 1) {
      try {
        Thread.sleep(1);
      } catch (InterruptedException ex) { println("Sleep interrupt: " + ex.getMessage()); }
    }
    return port.read();
  }
  
  // read until next newline character
  protected String readLine() {
    char b;
    String line = "";
    while ((b = (char)readByte()) != '\n') {
      line += b;
    }
    return line;
  }
  
  // read acknowledgement/error message
  protected void readAck() {
    int b = readByte();
    if (b != 0x01) {
      if (b == 0xF0) {
        // leftover init
        readByte();
      }
      if (b == 0xFF) {
        // error returned
        println("Error: " + this.readLine());
      }
    }
  }
  
  // reads until we receive the initialization magic bytes 0xFAF0
  protected void readInit() {
    int b = 0, lb = 0;
    while (b != 0xFA && lb != 0xF0) {
      // wait for magic bytes
      lb = b;
      b = readByte();
    }
  }
  
  /*
   * send messages
   */
   
   // send message and wait for acknowledgement
   protected void send(byte[] msg) {
    if (port == null)
        throw new RuntimeException("Serial device not connected");
     port.write(msg);
     readAck();
   }
   
   // initialize the driver with a given number of lights
   public void Setup(int lights) {
     this.lights = lights;
     byte[] msg = new byte[] {
       MSG_START,
       MSG_SETUP,
       1, // display id
       (byte)(3 * lights), // display width
       1, // display height
       1, // array width
       1, // array height
       MSG_END,
     };
     send(msg);
     println("Configured " + lights + " LEDs");
   }
   
   // flip display buffer command
   public void Flip() {
     byte[] msg = new byte[] {
       MSG_START,
       MSG_FLIP,
       MSG_END,
     };
     send(msg);
   }
  
   // set led strip pixels command
   public void Set(int start, color[] colors, boolean autoFlip) {
     final int headerlen = 5;
     
     byte[] msg = new byte[headerlen + 3 * colors.length + 1];
     msg[0] = MSG_START;
     msg[1] = MSG_SET;
     msg[2] = (byte)(autoFlip ? 1 : 0);
     msg[3] = (byte)start;
     msg[4] = (byte)colors.length;
     
     int i = headerlen;
     
     for (int c = 0; c < colors.length; c++) {
       // calculate gamma correct RGB
       int r = gamma_table[colors[c] >> 16 & 0xFF];
       int g = gamma_table[colors[c] >>  8 & 0xFF];
       int b = gamma_table[colors[c] >>  0 & 0xFF];
       
       msg[i++] = (byte)r;
       msg[i++] = (byte)g;
       msg[i++] = (byte)b;
     }
     
     msg[i] = MSG_END;
     send(msg);
   }
   
   // gamma correction lookup table
   private final int gamma_table[] = {
    0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
    0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  1,  1,  1,  1,
    1,  1,  1,  1,  1,  1,  1,  1,  1,  2,  2,  2,  2,  2,  2,  2,
    2,  3,  3,  3,  3,  3,  3,  3,  4,  4,  4,  4,  4,  5,  5,  5,
    5,  6,  6,  6,  6,  7,  7,  7,  7,  8,  8,  8,  9,  9,  9, 10,
   10, 10, 11, 11, 11, 12, 12, 13, 13, 13, 14, 14, 15, 15, 16, 16,
   17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 24, 24, 25,
   25, 26, 27, 27, 28, 29, 29, 30, 31, 32, 32, 33, 34, 35, 35, 36,
   37, 38, 39, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 50,
   51, 52, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 66, 67, 68,
   69, 70, 72, 73, 74, 75, 77, 78, 79, 81, 82, 83, 85, 86, 87, 89,
   90, 92, 93, 95, 96, 98, 99,101,102,104,105,107,109,110,112,114,
  115,117,119,120,122,124,126,127,129,131,133,135,137,138,140,142,
  144,146,148,150,152,154,156,158,160,162,164,167,169,171,173,175,
  177,180,182,184,186,189,191,193,196,198,200,203,205,208,210,213,
  215,218,220,223,225,228,231,233,236,239,241,244,247,249,252,255 };
}
