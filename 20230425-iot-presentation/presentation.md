---
marp: true
theme: "default"
backgroundImage: "linear-gradient(to bottom right, #fff, #ccc)"
color: "black"

author: "Daniel Schier"
title: "Manage IoT/Edge devices with Ansible + Podman"

paginate: true
header: "IoT"
footer: "Manage IoT/Edge devices with Ansible + Podman by @dschier"
---

## Manage IoT/Edge devices with Ansible + Podman

Very often, two simple tools are enough.

---

## Agenda

- 👋 Hi
- 📡 Internet of Things
- 🪴 My problem
- 🌊 Solution design
- 💾 Some code

---

## 👋 Hi

- My Name is Daniel.
- Some might say, I am a "DevOps engineer", ...
- ... but I see myself as an IT craftsman.
- I am also contributing to Open Source, democratizing software.
- I like cats, cookies, computer thingies (technical term), ...
- ... and Penguins, of course.

---

## 📡 Internet of Things

> "IoT is: Your toaster is selling NFT to pay the gambling debts of your fridge."
>
> *Somebody on the internet*

---

## 📡 Internet of Things

![bg right h:80%](./assets/iot_2.jpg)

Not really, but there are some challenges:

- IoT is not always online.
- IoT needs updates, too.
- IoT has some "special" security requirements.
- We might want to add/remove services.

---

## 🪴 My problem

First, let's talk about my problem:

- I don't want to learn 20 different tools
- I don't have budget for large infrastructure
- I don't want to publish my devices
- But, they should be manageable, reproducible, resilient, ...
- Documentation is important

---

## 🌊 Solution design

![bg right w:100%](./assets/iot_pull_only.drawio.png)

> Code
> \+ Hardware
> \+ Work
> \+ Code
> \___________________
> = Brrrr, Bing, Tada

---

## 🌊 Solution design

- ✅ Let's keep it simple
- ✅ Let's keep it cheap
- ✅ Don't experiment too much
- ✅ Be able to deploy services w/o touching the system

---

## 💾 Some "code"

`Code` incoming, but it's not much.

Find more at <https://github.com/dschier-wtd>

---

## 💾 Some "code" 1/3

Automate Configuration

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

## 💾 Some "code" 2/3

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

## 💾 Some "code" 3/3

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

## 🎬 Demo anyone?

Let's run the below:

```bash
/usr/bin/ansible-pull \
  --url https://github.com/dschier-wtd/fedora-iot.git \
  --inventory inventory/hosts.yml \
  --checkout main \
  ansible/playbooks/install_requirements.yml \
  ansible/playbooks/configure_server.yml
```

And see what happens.

---

## :heart: Thank you

You can find me all over the interwebs:

- 🐘 <https://fosstodon.org/@dschier>
- 🔧 <https://github.com/dschier-wtd>
- 🔧 <https://github.com/whiletruedoio>
- 🖋️ <https://blog.while-true-do.io>
- ✉️ dschier@while-true-do.io

Slides and example code at: <https://github.com/dschier-wtd/presentation/>.
