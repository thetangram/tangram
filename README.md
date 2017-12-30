# Tangram, an edge-side html composition server 

[![Build Status](https://travis-ci.org/thetangram/tangram.svg?branch=add_travis)](https://travis-ci.org/thetangram/tangram) [![Go Report Card](https://goreportcard.com/badge/github.com/thetangram/tangram)](https://goreportcard.com/report/github.com/thetangram/tangram) [![GoDoc](https://godoc.org/github.com/thetangram/tangram?status.svg)]


## Current status

**Tangram** is now in a conceptual phase. It's in scaffolding phase (configuring repo, setting up tooling, ...).


## Setup

### Requirements

  - [Make](https://www.gnu.org/software/make/) as build automation tool. 
  - [Docker](https://www.docker.com/) as container engine.


### Build

TL:DR

```
$ make 
``` 

This project uses ```make``` as build automation tool, so, all building project task are defined in ```Makefile```.

The project targets are:

  - ```compile``` (default target): Compiles the project.
  - ```dependencies```: Update project dependencies.
  - ```clean```: Clean all project artifacts
  - ```fmt```: Formats the code
  - ```test```: Run unit test
  - ```build``` (depends on ```fmt``` and ```test```): Compiles the project and generates a full independent Linux binary artifact. 
  - ```install``` (depends on ```build```): Pending. Build the Docker container image.
  - ```deploy``` (depends on ```install```): Pending. Publish the Docker container image to registry.

For example, to compile project while working: 

```
$ make 
``` 

To generate the final binary artifact:

```
$ make clean build 
``` 
