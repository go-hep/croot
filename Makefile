# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

ROOT_CONFIG := root-config
ROOT_CFLAGS := $(shell $(ROOT_CONFIG) --cflags)
ROOT_LDFLAGS := $(shell $(ROOT_CONFIG) --libs --ldflags) -lReflex -lCintex

CGO_LDFLAGS := "$(ROOT_LDFLAGS)"
CGO_CFLAGS  := "$(ROOT_CFLAGS) -Ibindings/inc -I."

# default to gc, but allow caller to override on command line
GO_COMPILER:=$(GC)
ifeq ($(GO_COMPILER),)
	GO_COMPILER:="gc"
endif

GO_VERBOSE := $(VERBOSE)
ifneq ($(GO_VERBOSE),)
	GO_VERBOSE:= -v -x
endif

# FIXME: until go-1.2 is released, we need to use 'goxx' instead of 'go'
#        so we can compile C++ files
GOCMD := goxx

build_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CPPFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) build $(GO_VERBOSE) -compiler=$(GO_COMPILER)

install_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CPPFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) install $(GO_VERBOSE) -compiler=$(GO_COMPILER)

test_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CPPFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) test $(GO_VERBOSE) -compiler=$(GO_COMPILER)

.PHONY: deps install

all: deps install

deps:
	@go get github.com/sbinet/goxx

install: deps
	@$(install_cmd) .
	@$(install_cmd) ./cmd/...

build: deps
	@$(build_cmd) .
	@$(build_cmd) ./cmd/...

test: install
	@$(test_cmd) .
