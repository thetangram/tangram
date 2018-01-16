build-image=jomoespe/go-build-base:1.0
version=$(shell cat VERSION)
user="$(shell id -u):$(shell id -g)"
build=$(shell git rev-parse --short HEAD)
buildDate=$(shell date --rfc-3339=seconds)
targetDir=/go/src/github.com/thetangram/tangram
docker-image=jomoespe/tangram
tangramd-name=tangramd

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

benchmark:
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _test-bench

#cover:
#	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _test-cover

build: clean fmt test
	@docker run --rm --user $(user) -v "$(PWD)":$(targetDir) -w $(targetDir) $(build-image) make _build

install: build
	@docker build --rm --build-arg version=$(version) -t "$(docker-image):$(version)" .
	@docker tag "$(docker-image):$(version)" "$(docker-image):latest"

deploy: install
	@docker push $(docker-image):$(version)
	@docker push $(docker-image):latest


_compile:
	@go build -v -o dist/$(tangramd-name) cmd/tangramd/main.go

_init-dependencies:
	@dep init -v

_dependencies:
	@dep ensure -update

_clean:
	@rm -rf dist/

_fmt:
	@go fmt ./...

_test:
	@go vet ./...
	@go test -cover ./...

_test-bench:
	@go test -bench=. ./...

#_test-cover:
#	@go test -coverprofile=coverage.out github.com/thetangram/tangram/pkg/...
#	@go tool cover -html=coverage.out -o cover.html

_build:
	@CGO_ENABLED=0 GOOS=linux go build -a \
									   -installsuffix cgo \
	                                   -ldflags "-s -w \
	                                             -X 'main.version=$(version)' \
	                                             -X 'main.build=$(build)' \
	                                             -X 'main.buildDate=$(buildDate)'"  \
	                                   -o dist/$(tangramd-name) \
									   cmd/tangramd/main.go
