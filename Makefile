# Go's default GOPATH doesn't include $PWD, so here's a lame workaround
GOPATH=$(shell pwd)

.PHONY: examples
examples: simple

simple: example/simple.go
	GOPATH=$(GOPATH) go build $<

clean:
	-rm simple
