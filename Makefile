# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

ROOT_CONFIG := root-config
ROOT_CPPFLAGS   := $(shell $(ROOT_CONFIG) --cflags)
ROOT_LDFLAGS  := $(shell $(ROOT_CONFIG) --libs --ldflags)

CGO_LDFLAGS := "$(ROOT_LDFLAGS)"
CGO_CFLAGS  := "$(ROOT_CFLAGS)"

# default to gc, but allow caller to override on command line
GO_COMPILER:=$(GC)
ifeq ($(GO_COMPILER),)
	GO_COMPILER:="gc"
endif

# FIXME: until go-1.2 is released, we need to use 'goxx' instead of 'go'
#        so we can compile C++ files
GOCMD := goxx

build_cwd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) build -compiler=$(GO_COMPILER) .

install_cwd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) install -compiler=$(GO_COMPILER) .

all: install

install:
	$(install_cwd)

build:
	$(build_cwd)
