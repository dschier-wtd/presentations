---
marp: true
theme: "rose-pine"

author: "Daniel Schier"
title: "DevOps Meetup Saxony"
header: "DevOps Meetup Saxony"
footer: "Let's pipeline © Daniel Schier 2024"

paginate: true
---

<div align="center">

![logo](./assets/logo.svg)

## Let's pipeline

DevOps and GitLab-CI

</div>

---

## :calendar: Agenda

Today, we will talk about:

1. :hugs: Hi
2. :carrot: Advertisements
3. :coconut: Dev(Sec)Ops Lifecycle
4. :mango: GitLab-CI
5. :broccoli: Let's pipeline (for real)

---

## :hugs: Hi

- My Name is Daniel.
- Some might say, I am a "DevOps engineer", ...
- ... but I see myself as an IT craftsman.
- I am also contributing to Open Source, democratizing software.
- I do some blogging, conferences and organizing meetups.
- I love cats, cooking, coffee, cookies, computer thingies (technical term), ...
- ... and Penguins, of course.

---

## :carrot: Advertisements

Have you heard of:

- DevOps Meetup
- DDOSUG
- Ansible Meetup
- Kubernetes and Cloud Native User Group
- DecompileD
- Chemnitzer Linuxtage

I need support for all of these!

---

## :coconut: Dev(Sec)Ops Lifecycle

![devops lifecycle](./assets/devsecops-lifecycle.png)

---

## :mango: GitLab-CI

![gitlab CI](./assets/gitlab-ci.png)

---

## :broccoli: Let's pipeline 1/8

User Story:

> As a user, I want to open my browser and point to a web site that replies
> with "Hello World!".

Solution Design:

> - modern language (Go)
> - should be in a DevOps Lifecycle
> - Continuous Integration for real!
> - Continuous Delivery for real!
> - Continuous Deployment for real!

---

## :corn: Let's pipeline 2/8

![bg right:30% contain](./assets/plan.png)

We want to use DevOps, right?

```yml
---
stages:

  - code
  - build
  - test
  - release
  - deploy
```

---

## :eggplant: Let's pipeline 3/8

![bg right:30% contain](./assets/code.png)

Is linting still part of coding? Maybe.

```yml

lint:
  stage: code
  image: docker.io/golangci/golangci-lint
  script: golangci-lint run -v

vet:
  stage: code
  image: docker.io/library/golang
  script: go vet ./...

```

---

## :banana: Let's pipeline 4/8

![bg right:30% contain](./assets/build.png)

Now, let's build the code.

```yml

build:
  stage: build
  image: docker.io/library/golang
  script: go build -v hello.go -o build/hello

```

---

## :zebra: Let's pipeline 5/8

![bg right:30% contain](./assets/test.png)

And run some tests.

```yml

unit_test:
  stage: test
  image: docker.io/library/golang
  script: go test -v -cover ./...
```

---

## :snail: Let's pipeline 6/8

![bg right:30% contain](./assets/release.png)

We want to upload our tested artifacts.

```yml
upload_package:
  stage: "release"
  image: "docker.io/library/python:latest"
  script:
    - 'curl --header "JOB-TOKEN: $CI_JOB_TOKEN"
      --upload-file build/hello
      $CI_API_V4_URL/projects/$CI_PROJECT_ID/packages/generic/$PACKAGE/$VERSION/hello'
```

---

## :peach: Let's pipeline 7/8

![bg right:30% contain](./assets/deploy.png)

And finally deploy it somewhere.

```yml
deploy_package:
  stage: "deploy"
  image: "docker.io/library/python:latest"
  script: "ansible-playbook ..."

```

---

## :heart: Thank you

You can find me all over the interwebs:

- 🐘 <https://fosstodon.org/@dschier>
- 🔧 <https://github.com/dschier-wtd>
- 🔧 <https://github.com/whiletruedoio>
- 🖋️ <https://blog.while-true-do.io>
- ✉️ [dschier@while-true-do.io]<mailto:dschier@while-true-do.io>

Slides and example code at: <https://github.com/dschier-wtd/presentation/>.
