from gpiozero import PWMLED
from signal import pause

led = PWMLED(4)

led.pulse()

pause()

