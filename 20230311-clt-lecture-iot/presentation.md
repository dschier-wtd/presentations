---
marp: true
theme: "default"
backgroundImage: "linear-gradient(to bottom right, #fff, #ccc)"
color: "black"

author: "Daniel Schier"
title: "Internet of Things (IoT) without Alexa, Google or else"

paginate: true
header: "CLT2023"
footer: "Internet of Things (IoT) without Alexa, Google or else by @dschier"
---

## Internet of Things (IoT) without Alexa, Google or else

Have you forgotten to water your plants, too?

---

## Agenda

- 👋 Hi
- 📡 Internet of Things
- 🪴 My problem
- 🌊 Solution design
- 🎸 Some hardware
- 🎵 Some software
- 🦾 Some work
- 💾 Some code

---

## 👋 Hi

- My Name is Daniel.
- Some might say, I am a "DevOps engineer".
- I am working for Wandelbots. We are democratizing robotics.
- I am also contributing to Open Source, democratizing software.
- I like cats, cookies, computer thingies (technical term), ...
- ... and Penguins, of course.

---

## 📡 Internet of Things

> "IoT is: Your toaster is selling NFT to pay the gambling debts of your fridge."
>
> *Somebody on the internet*

---

## 📡 Internet of Things 1/2

![bg right h:80%](./assets/iot_2.jpg)

But, technically ... :

- IoT device
- Edge
- Broker
- State-Machine
- Pull Principles
- Asynchronous Jobs
- Zigbee, LoRa, WiFi, Threat, ...

---

## 📡 Internet of Things 2/2

Use cases are:

- Home Automation
- Smart Grids
- Robots
- Public IP Webcams

---

## 🪴 My problem

![bg right h:80%](./assets/dead-plant.jpg)

---

## 🪴 My problem 1/2

First, let's talk about our problem:

- I want to have plants
- But... I always forget to water them
- Or water them too much

---

## 🪴 My problem 2/2

There are solutions, that:

- require an account somewhere
- want to send data somewhere
- cost a lot
- don't suite my need
- force me to use "a cool app"

---

## 🌊 Solution design

![bg right w:100%](./assets/solution_1.jpg)

> Code
> \+ Hardware
> \+ Work
> \___________________
> = Brrrr, Bing, Tada

---

## 🌊 Solution design

- ✅ Let's keep it simple
- ✅ Let's keep it cheap
- ✅ Don't experiment too much
- 🔄 Write in the Blog about it
- ❌ Get it done before CLT2023

---

## 🎸 Some hardware

![bg right h:80%](./assets/iot_2.jpg)

Welcome to the jungle.

---

## 🎸 Some hardware 1/5

We need to:

- measure the humidity
- send the measurements somewhere
- compute the recieved data
- send a notification

---

## 🎸 Some hardware 2/5

![bg right h:80%](./assets/raspberrypi.jpg)

For computing, almost everything works.

- must have network
- x86_64 (old Laptop, Home Server, ...)
- Raspberry Pi (or alternative)

---

## 🎸 Some hardware 3/5

![bg right h:80%](./assets/hygrometer.jpg)

Sensors is somewhat different.

- different sensors
- different outputs
- might work with ...

---

## 🎸 Some hardware 4/5

![bg right h:80%](./assets/esp32.jpg)

Processing Sensor Data ...

- Arduino
- ESP32
- ESP8266
- Battery or Cable?

Power? 80Mhz, 80KB RAM, 8MB Flash

---

## 🎸 Some hardware 5/5

My Bundle:

- Raspberry Pi 3B+
- Heltec WiFi Kit 8 (ESP8266)
- 3000 mAh battery
- AZDelivery Hygrometer

---

## 🎵 Some software

![bg right w:100%](./assets/logos.png)

Just another jungle.

---

## 🎵 Some software 1/7

![bg right w:150px](./assets/fedora_logo.png)

Operating System - Fedora IoT

- works well on Raspberry Pi
- immutable
- Podman included
- Greenboot

---

## 🎵 Some software 2/7

![bg right w:80%](./assets/podman_logo.png)

Containers - Podman

- Docker replacement
- daemonless
- rootless
- included in Fedora IoT
- GUI available
- auto-updates

---

## 🎵 Some software 3/7

![bg right w:80%](./assets/ansible_logo.png)

Automation - Ansible (Pull)

- in Package repository
- push and pull possible
- simple syntax
- Podman well supported

> Yes, there is an Ansible Community booth at CLT2023.

---

## 🎵 Some software 4/7

![bg right w:80%](./assets/nodered_logo.png)

State Maschine - Node Red

- supports MQTT
- supports Mail (and more)
- plug and pray

---

## 🎵 Some software 5/7

![bg right w:90%](./assets/mosquitto_logo.png)

Message Broker - Eclipse Mosquitto

- MQTT
- well maintained
- nice docs

---

## 🎵 Some software (optional) 6/7

![bg right w:60%](./assets/influxdb_logo.png)

Data Storing - InfluxDB

- supported from Node Red
- works well with time series data
- API implemented

---

## 🎵 Some software (optional) 7/7

![bg right w:60%](./assets/grafana_logo.png)

Data Display - Grafana

- Connectors to InfluxDB
- Powerful UI
- Lot's of prepared dashboards

---

## 🦾 Some work

We need to put some elbow grease in.

---

## 🦾 Some work 1/2

![bg right w:90%](./assets/battery.jpg)
![bg right w:90%](./assets/wifi_kit_8.jpg)

- Finding a battery solution.
- That does not need soldering.
- And has a display.

---

## 🦾 Some work 2/2

![bg right w:90%](./assets/cable.jpg)

- lots of cabeling.
- some soldering (optional)

---

## 💾 Some "code"

`Code` incoming, but it's not much.

Find more at <https://github.com/dschier-wtd>

---

## 💾 Some "code" 1/4

Automate Configuration

![bg right w:100%](./assets/iot_pull.jpg)

```yaml
---
- name: "Configure localhost"
  hosts: "localhost"

  tasks:

    # Ansible Pull
    - name: "Configure Ansible Pull Service"
      ansible.builtin.template:
        src: "ansible-pull.service.j2"
        dest: "/etc/systemd/system/ansible-pull.service"
        owner: "root"
        group: "root"
        mode: 0644
      become: true

...SNIP...

```

---

## 💾 Some "code" 2/4

Automate Container Deployment

```yaml
...SNIP...

    # container-web
    - name: "Download Image"
      containers.podman.podman_image:
        name: "docker.io/library/nginx"
        tag: "latest"
        state: "present"
      become: true
      tags:
        - "image"
        - "podman"
        - "web"

...SNIP...
```

---

## 💾 Some "code" 3/4

Automate Image Update

```yaml
...SNIP...

    - name: "Start & Enable Podman Auto Update Timer"
      ansible.builtin.service:
        name: "podman-auto-update.timer"
        state: "started"
        enabled: true
      become: true

...SNIP...
```

---

## 💾 Some "code" 4/4

Welcome to MQTT and NodeRed

![screenshot nodered](./assets/screenshot_nodered.png)

---

## :heart: Thank you

You can find me at the Ansible Booth and all over the interwebs:

- 🐘 <https://fosstodon.org/@dschier>
- 🔧 <https://github.com/dschier-wtd>
- 🖋️ <https://blog.while-true-do.io>
- ✉️ dschier@while-true-do.io

Slides and example code at: <https://github.com/dschier-wtd/presentations/>.
