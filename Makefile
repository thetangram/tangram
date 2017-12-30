build-image=jomoespe/go-build-base:1.0
version=$(shell cat VERSION)
user="$(shell id -u):$(shell id -g)"
build=$(shell git rev-parse --short HEAD)
buildDate=$(shell date --rfc-3339=seconds)
targetDir=/go/src/github.com/thetangram/tangram
docker-image=tangram:$(version)


compile:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _compile

clean:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _clean

init-dependencies:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _init-dependencies

dependencies:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _dependencies

fmt:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _fmt

test:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _test

build: fmt test
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _build

install: build
	@docker build --rm --build-arg version=$(version) -t $(docker-image) .

deploy: install
	@docker push docker-image


_compile:
	@go build -v

_init-dependencies:
	@dep init -v

_dependencies:
	@dep ensure -update

_clean:
	@go clean

_fmt:
	@go fmt ./...

_test:
	@go vet
	@go test


_build:
	@CGO_ENABLED=0 GOOS=linux go build -v \
	                                   -a -installsuffix cgo \
	                                   -ldflags "-s -w \
	                                             -X 'main.version=$(version)' \
	                                             -X 'main.build=$(build)' \
	                                             -X 'main.buildDate=$(buildDate)'" 
