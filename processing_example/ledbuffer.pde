public class LedBuffer 
{
  protected LedDriver driver;
  protected int lights;
  protected int ledsPerLight;
  protected color[] buffer;
  
  public LedBuffer(LedDriver driver, int lights, int ledsPerLight) {
    this.driver       = driver;
    this.lights       = lights;
    this.ledsPerLight = ledsPerLight;

    this.buffer       = new color[lights];
    this.driver.Setup(lights * ledsPerLight);
  }
  
  public color Get(int index) {
      return buffer[index];
  }

  public void Set(int index, color clr) {
    buffer[index] = clr;
  }
  
  public void Clear(color clr) {
    for(int i = 0; i < lights; i++) {
      buffer[i] = clr;
    }
  }
  
  public void Flip() {
    color[] frame = new color[lights * ledsPerLight];
    for(int i = 0; i < lights; i++) {
      for(int j = 0; j < ledsPerLight; j++) {
        frame[i * ledsPerLight + j] = buffer[i];
      }
    }
    driver.Set(0, frame, true);
  }
}
