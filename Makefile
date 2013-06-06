# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

ROOT_CONFIG := root-config
ROOT_CFLAGS := $(shell $(ROOT_CONFIG) --cflags)
ROOT_LDFLAGS := $(shell $(ROOT_CONFIG) --libs --ldflags) -lReflex -lCintex

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

INSTALL_DIR := $(firstword $(subst :, ,$(shell go env GOPATH)))/pkg/$(GOOS)_$(GOARCH)
INSTALL_LIBDIR := $(INSTALL_DIR)/lib

CGO_LDFLAGS := "-Wl,-rpath,$(INSTALL_LIBDIR) -L$(INSTALL_LIBDIR) -lcxx-croot"
CGO_CFLAGS  := "-Ibindings/inc -I."

CXX_CROOT_CXXFLAGS := $(ROOT_CFLAGS) -fPIC -Ibindings/inc -I.
CXX_CROOT_LDFLAGS := $(ROOT_LDFLAGS)

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
GOCMD := go

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

cxx_croot_sources := \
 bindings/src/croot.cxx \
 bindings/src/croot_class.cxx \
 bindings/src/croot_hist.cxx 

cxx_croot_objects := $(subst .cxx,.o,$(cxx_croot_sources))

.PHONY: deps install dirs clean

all: deps install

deps:
	@go get github.com/sbinet/goxx

dirs:
	@mkdir -p $(INSTALL_LIBDIR)

%.o: %.cxx
	$(CXX) $(CXX_CROOT_CXXFLAGS) -o $@ -c $<

install: deps cxx-lib
	@$(install_cmd) .
	@$(install_cmd) ./cmd/...

cxx-lib: dirs $(cxx_croot_objects)
	@$(CXX) -shared \
	 -o $(INSTALL_LIBDIR)/libcxx-croot.so \
	 $(CXX_CROOT_CXXFLAGS) $(CXX_CROOT_LDFLAGS) \
	 $(cxx_croot_objects)

test: install
	@$(test_cmd) .

clean:
	@rm -f $(cxx_croot_objects)
	@rm -f $(INSTALL_LIBDIR)/libcxx-croot.so
	@rm -f $(INSTALL_DIR)/github.com/go-hep/croot.a
