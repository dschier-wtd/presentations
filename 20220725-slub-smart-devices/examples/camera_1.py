from gpiozero import Button, LED
from picamera import PiCamera
from datetime import datetime
from signal import pause

button = Button(18)
led = LED(4)
camera = PiCamera()

def capture():
    timestamp = datetime.now().isoformat()
    camera.capture('/home/pi/%s.jpg' % timestamp)

def notify():
    print("Picture taken!")
    led.on

def picam():
    capture()
    notify()

button.when_pressed = picam
button.when_released = led.off

pause()
