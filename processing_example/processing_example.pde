// led color buffer
LedBuffer buffer;

void setup() {
    // device finder
    LedDeviceFinder devices = new LedDeviceFinder();

    LedDriver driver = new LedDriver(
        this, 
        devices.Guess()
    );

    buffer = new LedBuffer(
        driver, 
        LedConfig.LightCount, 
        LedConfig.LedsPerLight
    );
}

void draw() {
    int speed = 10;

    for(int x = 0; x < 256; x+=speed) {
        buffer.Clear(color(x, 0, 0, 255));
        buffer.Flip();
        delay(1);
    }
    for(int x = 0; x < 256; x+=speed) {
        buffer.Clear(color(255-x, 0, 0, 255));
        buffer.Flip();
        delay(1);
    }

    for(int x = 0; x < 256; x+=speed) {
        buffer.Clear(color(0, x, 0, 255));
        buffer.Flip();
        delay(1);
    }
    for(int x = 0; x < 256; x+=speed) {
        buffer.Clear(color(0, 255-x, 0, 255));
        buffer.Flip();
        delay(1);
    }

    for(int x = 0; x < 256; x+=speed) {
        buffer.Clear(color(0, 0, x, 255));
        buffer.Flip();
        delay(1);
    }
    for(int x = 0; x < 256; x+=speed) {
        buffer.Clear(color(0, 0, 255-x, 255));
        buffer.Flip();
        delay(1);
    }
}