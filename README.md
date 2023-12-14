<!--
reference: https://www.makeareadme.com/
reference: https://commonmark.org/
-->

[![Cirrus CI - Base Branch Build Status](https://img.shields.io/cirrus/github/dschier-wtd/fedora-workstation?logo=Cirrus-ci)](https://cirrus-ci.com/github/dschier-wtd/presentations)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dschier-wtd/fedora-workstation?logo=GitHub&label=Release&sort=semver)](https://github.com/dschier-wtd/presentations/releases)
[![GitHub issues](https://img.shields.io/github/issues/dschier-wtd/fedora-workstation)](https://github.com/dschier-wtd/presentations/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/dschier-wtd/fedora-workstation)](https://github.com/dschier-wtd/presentations/pulls)
[![GitHub license](https://img.shields.io/github/license/dschier-wtd/fedora-workstation)](https://github.com/dschier-wtd/presentations/blob/main/LICENSE)

# Presentations

Material for presentations and talks I have given.

## Motivation

I just wanted to publish my presentation content that I worked on for
conferences, meetups and more. Maybe it is helpful for someone else, too.

## Description

The repository contains all kinds of unsorted material like presentations, code
examples, snippets and more. Furthermore, I have added a couple of links to
other repositories, where you can find presentations from me.

## Content

I am trying to gather my presentation content in this repository now. Since my
slides and example code is always in **English** language, it may be useful for
you. Most recordings are done in German language, though.

### 2023

#### DevOps Meetup - Enterprise Linux

Giving an overview about Enterprise Linux like AlmaLinux OS, Rocky Linux or
RHEL and Ubuntu LTS.

- event: DevOps Saxony Meetup
- type: overview/impulse talk
- date: 12.12.2023
- language: english/german
- links: [slides & code](./20231212-enterprise-linux/)

#### Tux Tage - DevOps Ideals

A talk for the [Tux Tage 2023](https://www.tux-tage.de/) about DevOps and
how can have a local development, build, release environment.

- event: [Tux Tage 2023](https://www.tux-tage.de/)
- type: guide/tutorial
- date: 12.11.2023
- language: english/german
- links: [slides & code](./20231112-devops-simplicity/)

#### Ansible - The hard way/things

A talk for the Ansible Community Day Berlin 2023

- event: [Ansible Community Day](https://www.ansible.com/blog/ansible-community-day-berlin-2023)
- type: short talk
- date: 20.09.2023
- language: english
- links: [presentation](./20230920-ansible-hard-things/)

#### DevOps Meetup - Types and Kinds of Monitoring

A short intro talk to our DevOps Saxony meetup about monitoring.

- event: DevOps Saxony meetup
- type: short talk
- date: 13.09.2023
- language: english
- links: [presentation](./20230913-types-of-monitoring/)

#### Cloud Architecture Discussion

Provided an example architecture for a cloud hosted checkmk environment.

- event: Ad-Hoc presentation
- type: short talk/open discussion
- date: 17.08.2023
- language: english
- links: [presentation](./20230817-cloud-architectures/)

#### DevOps Ideals - Simplicity & Locality

A practical guide to apply simplicity and locality to your code.

- event: Ad-Hoc presentation
- type: short talk
- date: 27.06.2023
- language: english
- links: [demo, code and presentation](./20230627-devops-simplicity/)

#### Ansible - IoT

A short talk about IoT with Ansible and Podman.

- event: Ad-Hoc presentation
- type: Lightning Talk
- date: 25.04.2023
- language: english
- links: [demo code and slides](./20230425-iot-presentation/)

#### Ansible - Dynamic Inventories

A short lightning talk for the Ansible Meetup.

- event: Ansible Meetup Dresden
- type: Lightning Talk
- date: 13.04.2023
- language: german
- links: [demo code](./20230413-ansible-meetup/)

#### Internet of Things (IoT) without Alexa, Google or else

A lecture for the Chemnitz Linux Days/Chemnitzer Linuxtage, explaining how
one can do IoT, without using Google, Alexa or other cloud providers.

- event: CLT 2023
- type: Lecture
- date: 11.03.2023
- language: english/german
- links: [docs and examples](./20230311-clt-lecture-iot)

#### DevOps Workshop

A workshop to explain DevOps ideals, metrics and KPIs. It was also about
decision methods and frameworks. Team topologies was in there and some ideas
to design DevSecOps pipelines.

- event: DevOps Workshop
- type: Workshop
- date: 19.01.2023
- language: english
- links: [docs and examples](./20230119-devops-workshop/)

### 2022

#### Smart Devices

A two part workshop about programming smart devices with a Raspberry Pi. It is
mostly about starter examples for coding, but also about starting the Raspberry
with an attached display in a Kiosk mode.

At the end, one is able to use gpiozero and Python on a Raspberry to turn on
an LED via touchscreen, web, button and command line. A good starting point for
further investigation and experimenting.

- event: SLUB Smart Displays / Smart Devices
- type: Workshop
- date: 25.07.2022, 01.08.2022
- language: english
- links 1: [slides and examples](./20220725-slub-smart-devices-1/)
- links 2: [slides and examples](./20220801-slub-smart-devices-2/)

#### The perfect SCRUM sprint

This was a BarCamp style session and the presentation material is more or less
for orientation purposes. It wasn't used during the session but to prepare my
course and get an idea what I want to talk about. The attached drawio file was
drawn to the board and we talked about optimization potentioals to take care of
work weeks, avoid weekend work and such stuff.

- event: Agile Meetup #30
- type: Lecture
- date: 29.03.2022
- language: english
- links: [slides](./20220329-agile-dresden-meetup/perfect_scrum.pdf),
  [drawio](./20220329-agile-dresden-meetup/perfect_scrum.drawio)

#### Ansible - The very first steps

A 30 minute talk about the very first steps to get started with Ansible. It was
about the installation, basic conceots and writing the first mini playbook.

- event: Chemnitzer Linuxtage 2022
- type: Lecture
- date: 12.03.2022
- language: english
- links: [link](https://chemnitzer.linux-tage.de/2022/en/programm/beitrag/155),
  [slides](./20220312-clt-ansible/ansible-first-steps.pdf),
  [code](./20220312-clt-ansible/ansible-first-steps/)

#### Ansible in a Cloud (Native) World

The talk was about how Ansible can be used in a Cloud (Native) world and how it
can help small teams. Avoiding the "learn a new the tool everyday"-loop and
one-trick-ponies were also part of the talk. In the demo part, a k3s setup and
deployment of a simple Kubernetes application was shown.

- event: Chemnitzer Linuxtage 2022
- type: Lecture
- date: 12.03.2022
- language: german/english
- links: [link](https://chemnitzer.linux-tage.de/2022/en/programm/beitrag/156),
  [video](https://chemnitzer.linux-tage.de/2022/en/programm/beitrag/156),
  [slides](./20220312-clt-ansible/ansible-cloud-native.pdf),
  [code](./20220312-clt-ansible/ansible-cloud-native/)

#### Ansible 2.12

A talk for the Ansible community presentation booth at the CLT showing new
features and changes since Ansible 2.9.

- event: Chemnitzer Linuxtage 2022
- type: Talk
- date: 12.03.2022
- language: english
- links: [link](https://chemnitzer.linux-tage.de/2022/en/programm/beitrag/155),
  [slides](./20220312-clt-ansible/ansible-2.12.pdf)

#### Fedora Smart Display and Mirror

A talk for the "Makers' March Meetup" to demonstrate my Fedora Smartdisplay
code and talk about the setup. There will be Raspberries and Ansible.

- event: "Makers' March Meetup (DDOSUG + SLUB Makerspace)"
- type: Talk
- date: 10.03.2022
- language: german
- links: [slides](./20220310-makers-march-meetup/makers_march_smart_display.pdf),
  [code](https://github.com/dschier-wtd/fedora-smartdisplay/)

#### Ansible Validation and Integration Testing

A lecture about testing in Ansible and demonstrating how argument specifications
work. It is also shown how integration testing can be done in collections.

- event: Ansible Anwendertreffen 2022
- type: Lecture
- date: 15.02.2022
- language: german/english
- links: [link](https://www.ansible-anwender.de/post/2022/02/rueckblick/),
  [video](https://www.youtube.com/watch?v=Cj_d1N8MIEM),
  [slides](./20220215-ansible-anwendertreffen/Ansible_Validation_Integration_Testing_Roles.pdf),
  [code](./20220215-ansible-anwendertreffen/ansible/)

### 2021

#### Immutable Linux

A short talk about different immutable GNU/Linux distributions.

- event: DDOSUG Meetup 12/2021
- type: talk
- date: 15.12.2021
- language: german/english
- links: [video](https://www.youtube.com/watch?v=HGNs6qBWlXY),
  [slides](./20211215-ddosug-meetup/immutable-linux.pdf)

#### Ansible + Podman for IoT and Tinkering

A lecture explaining how Ansible can be used in IoT situations and in
combination with Podman to provide an automatic update pull behavior for
configuration and software updates.

- event: Ansible Anwendertreffen 2021
- type: Lecture
- date: 18.05.2021
- language: german/english
- links: [link](https://www.ansible-anwender.de/post/2021/05/rueckblick/),
  [video](https://www.youtube.com/watch?v=ZsFLzEgK-_w),
  [slides](./20210518-ansible-anwendertreffen/ansible_podman_for_iot_edge_and_tinkering.pdf),
  [code](./20210518-ansible-anwendertreffen/)

#### Micro Editor

A talk about the micro editor and how it compares to nano and vim.

- event: DDOSUG Meetup 05/2021
- type: talk
- language: german/english
- links: [video](https://www.youtube.com/watch?v=rGQj9W4ho2Y)

#### Ansible - Wiederverwendbarer Code (ger)

This talk is part of a webinar series for my employer profi.com AG that shows
how you can start developing re-usable code in roles and collections.

- event: profi.com AG Webinar series
- type: webinar
- language: german
- links: [video](https://www.youtube.com/watch?v=ANPx8hFXqog)

#### Ansible Molecule with Podman/Docker

A talk about Testing with Molecule and facilitating Podman and Docker for
integration and functional testing of Ansible Roles.

- event: DDOSUG Meetup 04/2021
- type: talk
- language: german/english
- links: [video](https://www.youtube.com/watch?v=W5xLYmsa9uk)

#### Ansible - Ein Überblick

A webinar to start a series about Ansible for my employer profi.com AG. In the
talk I am explaining how the ecosystem of Ansible looks like and what you can
expect if you want to dig into Ansible.

- event: profi.com AG Webinar series
- type: webinar
- language: german
- links: [video](https://www.youtube.com/watch?v=AiMjFiS-NY8)

#### Vagrant + ein bisschen Ansible

A talk showing how Vagrant works and how you can use Ansible to configure the
vagrant machines.

- event: DDOSUG Meetup 03/2021
- type: talk
- language: german/english
- links: [video](https://www.youtube.com/watch?v=U6Km6_H85Fk)

#### Ansible Collections - Das "Was", "Warum" und "Wie"

A short talk about Ansible collections, that demonstrates how collections work
and how you can use them.

- event: DDOSUG Meetup 01/2021
- type: talk
- language: german/english
- links: [video](https://www.youtube.com/watch?v=tFUQ065obXU)

### 2020

TBD

### 2019

TBD

- [CLT 2019: Setup your workstation with Ansible (ger)](https://chemnitzer.linux-tage.de/2019/de/programm/beitrag/268)

## Contribute

The repository contains content and things that happened in the past. Therefore,
I am not investing much time to keep code examples up-to-date or fiddle with
minor issues. Nevertheles, if you find something, that is either downright
wrong or content is not usable from your side, please feel free to open an
issue or provide a pull request.

## License

Except otherwise noted, all work is [licensed](LICENSE) under a
[BSD-3-Clause License](https://opensource.org/licenses/BSD-3-Clause).

## Contact

Please feel free to reach out to me to provide feedback or get in touch.

- Site: <https://while-true-do.io>
- Blog: <https://blog.while-true-do.io>
- Code: <https://github.com/dschier-wtd>
- Mail: [dschier@while-true-do.io](mailto:dschier@while-true-do.io)
