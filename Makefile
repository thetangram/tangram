build-image=jomoespe/go-build-base:1.0
version=$(shell cat VERSION)
user="$(shell id -u):$(shell id -g)"
build=$(shell git rev-parse --short HEAD)
buildDate=$(shell date --rfc-3339=seconds)


compile:
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _compile

clean:
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _clean

fmt:
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _fmt

test:
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _test

build: fmt test
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _build

install: build
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _install

deploy: install
	@docker run --rm --user $(user) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(build-image) make _deploy


_compile:
	@go build -v

_clean:
	@go clean

_fmt:
	@go fmt

_build:
	@CGO_ENABLED=0 GOOS=linux go build -v \
	                                   -a -installsuffix cgo \
	                                   -ldflags "-s -w \
	                                             -X 'main.version=$(version)' \
	                                             -X 'main.build=$(build)' \
	                                             -X 'main.buildDate=$(buildDate)'" 

_test:
    @echo "Pending: here run the unit test"

_install:
    @echo "Pending: here the build the docker image"

_deploy:
    @echo "Pending: here deploy (publish) the docker image"
