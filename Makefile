# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

ROOT_CONFIG := root-config
ROOT_CFLAGS := $(shell $(ROOT_CONFIG) --cflags)
ROOT_VERSION := $(shell $(ROOT_CONFIG) --version | cut -f1 -d.)
ifeq ($(ROOT_VERSION),6)
ROOT_LDFLAGS := $(shell $(ROOT_CONFIG) --libs --ldflags)
gocroot_tag := "root6"
gendict_file := gen-goedm-dict-root6.go
else
gocroot_tag := "root5"
ROOT_LDFLAGS := $(shell $(ROOT_CONFIG) --libs --ldflags) -lReflex -lCintex
gendict_file := gen-goedm-dict-root5.go
endif

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

INSTALL_DIR := $(firstword $(subst :, ,$(shell go env GOPATH)))/pkg/$(GOOS)_$(GOARCH)
INSTALL_LIBDIR := $(INSTALL_DIR)/github.com/go-hep/croot/_lib

ifeq ($(GOOS),linux)
CGO_LDFLAGS := "-Wl,-rpath,$(INSTALL_LIBDIR) -L$(INSTALL_LIBDIR) -lcxx-croot $(ROOT_LDFLAGS)"
endif
ifeq ($(GOOS),darwin)
CGO_LDFLAGS := "-L$(INSTALL_LIBDIR) -lcxx-croot $(ROOT_LDFLAGS)"
endif
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

GOCMD := go

gen_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) generate $(GO_VERBOSE) -tags=$(gocroot_tag) -compiler=$(GO_COMPILER)

build_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) build $(GO_VERBOSE) -tags=$(gocroot_tag) -compiler=$(GO_COMPILER)

install_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) install $(GO_VERBOSE) -tags=$(gocroot_tag) -compiler=$(GO_COMPILER)

test_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) test $(GO_VERBOSE) -tags=$(gocroot_tag) -compiler=$(GO_COMPILER)

clean_cmd = \
 CGO_LDFLAGS=$(CGO_LDFLAGS) \
 CGO_CFLAGS=$(CGO_CFLAGS) \
 $(GOCMD) clean $(GO_VERBOSE) -tags=$(gocroot_tag) -compiler=$(GO_COMPILER)

cxx_croot_sources := \
 bindings/src/croot.cxx \
 bindings/src/croot_class.cxx \
 bindings/src/croot_go_schema.cxx \
 bindings/src/croot_goobject.cxx \
 bindings/src/croot_hist.cxx \
 bindings/src/croot_interpreter.cxx \
 bindings/src/croot_leaf.cxx

cxx_croot_dicts := \
 bindings/src/goedm_dict.cxx

cxx_croot_objects := $(subst .cxx,.o,$(cxx_croot_sources))
cxx_croot_dict_objects := $(subst .cxx,.o,$(cxx_croot_dicts))
cxx_croot_dict_pch := bindings/src/goedm_dict_rdict.pcm

.PHONY: install dirs clean

all: install

dirs:
	@mkdir -p $(INSTALL_LIBDIR)

%.o: %.cxx
	@$(CXX) $(CXX_CROOT_CXXFLAGS) -o $@ -c $<

install: cxx-lib
	@$(gen_cmd) .
	@$(install_cmd) .
	@$(install_cmd) ./cmd/...

cxx-lib: dirs $(cxx_croot_objects) $(cxx_croot_dict_objects)
	@$(CXX) -shared \
	 -o $(INSTALL_LIBDIR)/libcxx-croot.so \
	 $(CXX_CROOT_CXXFLAGS) $(CXX_CROOT_LDFLAGS) \
	 $(cxx_croot_objects) \
	 $(cxx_croot_dict_objects)
	@touch root.go
	@install -m644 \
		./$(cxx_croot_dict_pch) \
		$(INSTALL_LIBDIR)/$(notdir $(cxx_croot_dict_pch))

gen-dict:
	@cd bindings/src && $(gen_cmd) $(gendict_file)

$(cxx_croot_dicts): gen-dict


test: install
	@$(test_cmd) ./cmem
	@$(test_cmd) ./cgentype
	@$(test_cmd) .

clean:
	@rm -f object_impl.go
	@rm -f $(cxx_croot_objects)
	@rm -f $(cxx_croot_dicts) $(cxx_croot_dict_objects) $(cxx_croot_dict_pch)
	@rm -rf $(INSTALL_LIBDIR)
	@rm -rf $(INSTALL_DIR)/github.com/go-hep/croot
	@rm -rf $(INSTALL_DIR)/github.com/go-hep/croot.a
	@$(clean_cmd) ./...
