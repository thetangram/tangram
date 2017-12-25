BUILD-IMAGE=jomoespe/go-build-base:1.0

build:
	@docker run --rm --user $(shell id -u):$(shell id -g) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(BUILD-IMAGE) make _build 

clean:
	@docker run --rm --user $(shell id -u):$(shell id -g) -v "$(PWD)":"$(PWD)" -w "$(PWD)" $(BUILD-IMAGE) make _clean 


_clean:
	@go clean

_build:
	@go build -v

