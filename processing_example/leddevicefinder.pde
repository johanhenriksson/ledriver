import processing.serial.*;

class LedDeviceFinder 
{
    public String[] List() {
        String[] devices = Serial.list();
        return devices;
    }

    public String Select(int index) {
        String[] devices = List();
        return devices[index];
    }

    public void PrintList() {
        String[] devices = List();
        for(int i = 0; i < devices.length; i++) {
            print(i);
            print(": ");
            print(devices[i]);
            print("\n");
        }
    }

    public String Guess() {
        String[] devices = List();
        for(String device : devices) {
            if (
                device.contains("tty.usbmodem") || // OSX
                device.contains("ttyACM")          // Linux
            ) {
                return device;
            }
        }
        throw new RuntimeException("No serial device found");
    }
}
