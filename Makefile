# Comments starting with #: below are remake GNU Makefile comments. See
# https://github.com/rocky/remake/wiki/Rake-tasks-for-gnu-make
#
# Supplying a make file prevents travis from trying to build
# using go get -t -v ../... which tries to build
# demo code.

.PHONY: all install check test

#: Same as go test
all: test

#: Just build
build: columnize.go
	go build

# Same as build
check: build

#: Install this
install: columnize.go test
	go install

#: The GNU Readline REPL front-end to the go-interactive evaluator
test: columnize.go columnize_test.go
	go test
