from gpiozero import LED
from signal import pause

led = LED(4)

led.blink()

pause()
