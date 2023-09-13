# Architecture

This document describes an architecture proposal for further prototyping, proof
of concept and decision making for the checkmk-SaaS-service.

## Motivation

We want to offer the checkmk as a hosted software as a service (SaaS). A cloud
provider is not chosen, yet. A document is required for evaluation and further
prototyping.

In the end, we want to make a decision that supports our requirements and common
demands on a cloud service.

## Requirements

The below requirements were given with different importance.

### Now

- first beta in 3 months
- integration of standard service
- service provides a http(s) interface
- state management w/o database
- common security best practices
- segregation of development and operations

### Next

- backup & disaster recovery
- updates
- maintenance
- monitoring and alerting

### Later

- multi-tenant for customers with different criticality and security demands

## Architecture

The following sections will provide a description of the overall architecture
for a web based service. In this case, the checkmk raw (community) edition
is chosen for deployment.

### Landing Zones

For now, I focused on the public landing zone, which has the highest demand
regarding security and availability.

![Landing Zones](./assets/landing_zones.drawio.png)

As the entrance to the entire infrastructure a Web Application Firewall should
be installed. This allows fine-tuning for most demands. Most cloud providers
also combine the WAF functionality with a CDN option and this makes it easier
to provide regional sharding later on.

As a second separation one should ensure that only the desired traffic is
reaching the compute nodes. This can be achieved in multiple ways. The most
common options are NAT gateways, Network- or Application Loadbalancers and
dedicated Ingress nodes (for Kubernetes only). The final setup should be
considered based on the cloud provider decision. Pricing and features vary a
lot, here.

In a dedicated private network, the real workload will be done.

### Containerization

A containerized solution was considered as the "best" option, due to the
following reasons:

- distribution independent
- cloud provider independent
- multiple container engines available (kubernetes, podman, docker, etc.)
- easy to reproduce locally
- sand-boxing can be added
- easy to automate
- container images already existing
- well supported in all major hyperscaler

But it also introduces some challenges:

- stateful applications need volumes and therefore some kind of block- or
  network storage
- conflicting ports can be problematic w/o ingress handling

### Kubernetes

As a containerization technology, I opted for Kubernetes. It is very well known
for management of scalable container applications. All cloud providers offer
a wide variety of Kubernetes workload services. The below diagram should give
a brief overview how such a setup may look like.

![Kubernetes](./assets/kubernetes.drawio.png)

As you can see, this will require additional separation to ensure that the
control plane is not directly reachable via customer facing interfaces.

### Single Instances

Each single deployment will be considered to be a container, for now. The
demand for very huge instances is not forseeable in the first beta iteration.
Therefore, each customer can access a vertical scalable container instance in
his own namespace.

![single instance](./assets/single_instance.drawio.png)

As you can see, we will use a block storage via a dedicated Persistant Volume
Claim. This is supported by most cloud providers via dedicated storage drivers.

### Automation (DRAFT)

The automation should be separated to a control cluster to foster security
and avoid resource conflicts.

### Updates (DRAFT)

Updates of single instances can be triggered in the related repositories, which
will trigger an automated deployment of newer versions. to ensure that these
updates are working well, canary deployment scenarios, multi staging and
rollbacks should be considered.

### Backup & Disaster Recovery (DRAFT)

The only known stateful resource is located in the block storage. A backup of
the same is supported by all major cloud providers.

### Monitoring & Alerting (DRAFT)

- eat-your-own-dogfood vs cncf

### Security Concepts (DRAFT)

## Prototyping (DRAFT)

### Pre-requisites

For the initial phase, the following pre-requisites and assumptions were taken.

- agents pushing metrics to our instance via HTTPS
- each customer has a dedicated instance
- instances are stateful
- public APIs are consumed actively (polling/pulling)

### Implementation (DRAFT)

## Risk Assessment (DRAFT)

