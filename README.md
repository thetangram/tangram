Tangram
=======

Requirements
------------

  - [Make](https://www.gnu.org/software/make/) as build automation tool. 
  - [Docker](https://www.docker.com/) as container engine.


How to build
------------

TL:DR

```
$ make 
``` 

This project uses **make** as build automation tool, so, all building project task are driven by **make**.

The project Makefile have this targets:

  - **compile** (default target): Compiles the project.
  - **clean**: Clean all project artifacts
  - **fmt**: Formats the code
  - **test**: Run unit test
  - **build** (depends on fmt and test): Compiles the project and generates a full independent Linux binary artifact. 
  - **install** (depends on build): Pending. Build the Docker container image.
  - **deploy** (depends on install): Pending. Publish the Docker container image to registry.

For example, to compile project while working: 

```
$ make 
``` 

To generate the final binary artifact:

```
$ make clean build 
``` 
