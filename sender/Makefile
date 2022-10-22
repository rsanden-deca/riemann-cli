SHELL:=/bin/bash

.PHONY: all clean exe test

all: exe

clean:
	rm -rf tmp

exe: clean
	go build

test:
	go test
