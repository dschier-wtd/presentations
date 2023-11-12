---
marp: true
theme: "default"
backgroundImage: "linear-gradient(to bottom right, #fff, #ddd)"
color: "black"
style: |
  pre {
   font-size: 12px;
  }

author: "Daniel Schier"
title: "DevOps Ideals - Simplicity"

paginate: true
---

## 👨‍💻 DevOps Ideals

> Simplicity & Locality

A practical guide to simpler workflows and local build environments.

---

## 📋 Agenda

- 👋 Hi
- 📡 DevOps Ideals
- 🪴 Simplicity & Locality
- 🌊 A practical example
- 🎬 Demo & Code
- 💡 Learnings

---

## 👋 Hi

- My Name is Daniel.
- Some might say, I am a "DevOps engineer", ...
- ... but I see myself as an IT craftsman.
- I am also contributing to Open Source, democratizing software.
- I do some blogging, conferences and organizing meetups.
- I love cats, cooking, coffee, cookies, computer thingies (technical term), ...
- ... and Penguins, of course.

---

## 📡 DevOps Ideals 1/2

![bg right w:100%](./assets/bulb.drawio.png)

Something to aim for and never reach.

---

## 📡 DevOps Ideals 2/2

- Simplicity & Locality
- Focus, Flow & Joy
- Improvement of daily work
- Psychological safety
- customer focus

---

## 🪴 Simplicity & Locality 1/2

![drop-shadow width:1400px](./assets/devops_ideals_1.drawio.png)

---

## 🪴 Simplicity & Locality 2/2

The ideal of locality and simplicity states, that every workflow should be
simple, easy to follow and, in the best case, reproducible by a single person.

> This can be achieved by:
>
> - simple applications
> - build in a Starbucks coffee shop in Madrid
> - don't worry about complicated/central tooling
> - stay in the repository

---

## 🌊 A practical example

Story Description:

> As a user, I love to be greeted. Therefore, I want to be greeted from an API
> with a "Hello World!".

Acceptance Criteria:

- new service
- following simplicity & locality
- can be understood by the PM
- functional testing provided
- deployable to Kubernetes (ingresss, service and workload)

---

## 🎬 Demo & Code 1/8

`Code` incoming, but it's not much.

Find more at <https://github.com/dschier-wtd/presentations/>

---

## 🎬 Demo & Code 2/8

The code (in Golang)

```golang
package main

import (
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {

  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  e.Get("/health", func(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
  })

  e.Logger.Fatal(e.Start(":3000"))
}
```

We can build this locally:

```shell
go build .
```

---

## 🎬 Demo & Code 3/8

The build (Containers)

```Dockerfile
FROM docker.io/library/golang:1.20 AS build

WORKDIR /usr/local/src
ENV CGO_ENABLED=0

COPY src/ /usr/local/src/

RUN go mod tidy \
    && go get . \
    && go build -o /usr/local/bin/hello -v .

FROM scratch

COPY --from=build /usr/local/bin/hello /usr/local/bin/

EXPOSE 3000/tcp

CMD ["/usr/local/bin/hello"]
```

We can build this locally:

```shell
podman build -t go-hello .
```

---

## 🎬 Demo & Code 4/8

The unit test (go)

```shell
# Run unit tests
$ go test -v ./...

# Maybe with coverage?
$ go test -v -cover ./...
```

Go has some awesome built-in capabilities.

---

## 🎬 Demo & Code 5/8

The integration test (Hurl)

```hurl
GET http://{{host}}/
GET http://{{host}}/health

HTTP 200
```

We can run this locally:

```shell
# Start
podman run -p 3000:3000 localhost/go-hello

# Smoke
curl localhost:3000
"Hello, World!"

# Functional
hurl --variable host=localhost:3000 tests/functional.hurl
```

---

## 🎬 Demo & Code 6/8

The release (local registry)

```shell
# Tag properly
podman image tag localhost/go-hello docker.io/dschier/go-hello:latest

# Push to registry
podman image push podman image push docker.io/dschier/go-hello:latest
```

---

## 🎬 Demo & Code 7/8

The deployment (Minikube)

```shell
$ tree kubernetes/
kubernetes/
├── deployment.yaml
├── ingress.yaml
├── namespace.yaml
└── service.yaml

1 directory, 4 files
```

We can start this locally again:

```shell
# Deploy
kubectl apply -f kubernetes/

# Test it again
url 192.168.49.2
hurl --variable host=192.168.49.2 tests/functional.hurl
```

---

## 🎬 Demo & Code 8/8

The documentation (mkdocs)

```shell
# Python venv
python -m venv .venv
source .venv/bin/activate

# Install deps
pip install -r requirements.txt

# Build docs
mkdocs serve
```

---

## 💡 Improvements to be made

Possible Improvements?

- multi stage deployments
- scripted build/test/release
- addition of a CI/CD process
- package kubernetes (Helm)
- security scans
- way more docs
- unit tests for src/

---

## 💡 Summary

Ok, what have we done?

> Plan, Code, Build, Test, Release, Deploy

- we followed a software development lifecycle from Plan to Deploy
- we can code, build, test, release & deploy locally
- Go is super simple for small apps
- Containers are awesome
- Shown us lots of things we already knew

---

## ❤️ Thank you

You can find me all over the interwebs:

- 🐘 <https://fosstodon.org/@dschier>
- 🔧 <https://github.com/dschier-wtd>
- 🔧 <https://github.com/whiletruedoio>
- 🖋️ <https://blog.while-true-do.io>
- ✉️ [dschier@while-true-do.io]<mailto:dschier@while-true-do.io>

Slides and example code at: <https://github.com/dschier-wtd/presentation/>.
