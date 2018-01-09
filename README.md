# Tangram, an edge-side html composition server 

[![Build Status](https://travis-ci.org/thetangram/tangram.svg?branch=add_travis)](https://travis-ci.org/thetangram/tangram) [![Go Report Card](https://goreportcard.com/badge/github.com/thetangram/tangram)](https://goreportcard.com/report/github.com/thetangram/tangram) [![Coverage](http://gocover.io/_badge/github.com/thetangram/tangram)](http://gocover.io/github.com/thetangram/tangram)  [![GoDoc](https://godoc.org/github.com/thetangram/tangram?status.svg)](https://godoc.org/github.com/thetangram/tangram)


## Current status

**Tangram** is now in a conceptual phase. It's in scaffolding phase (configuring repo, setting up tooling, ...).


## Setup

### Requirements

  - [Make](https://www.gnu.org/software/make/) as build automation tool. 
  - [Docker](https://www.docker.com/) as container engine.


### Build & Run

TL:DR

```
$ make build 
$ ./dist/tangram
``` 

#### Building

This project uses ```make``` as build automation tool, so, all building project task are defined in ```Makefile```.

The project targets are:

  - ```compile``` (default target): Compiles the project.
  - ```dependencies```: Update project dependencies.
  - ```clean```: Clean all project artifacts
  - ```fmt```: Formats the code
  - ```test```: Runs unit tests
  - ```benchmark```: Runs benchmark tests
  - ```build``` (depends on ```fmt``` and ```test```): Compiles the project and generates a full independent Linux binary artifact. 
  - ```install``` (depends on ```build```): Build the Docker container image.
  - ```deploy``` (depends on ```install```): Publish the Docker container image to registry.

For example, to compile project while working: 

```
$ make 
``` 

To generate the final Linux full independent binary artifact:

```
$ make clean build 
``` 

Also you can generate a Docker image:

```
$ make install 
``` 


#### Running

You can run the binary artifact or the Docker container.

To run the binary:

```
$ ./dist/tangram
``` 

To run the Docker image:

```
$ docker run -ti -p 2018:2018 tangram
```

The service exposes *liveness* and *readiness* health checks. You can access from command line (```curl http://localhost:2018/healthy``` or ```curl http://localhost:2018/ready```) of from a browser (```xdg-open http://localhost:2018/healthy``` or ```xdg-open http://localhost:2018/ready```).
