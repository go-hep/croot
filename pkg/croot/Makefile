# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=bitbucket.org/binet/go-croot/pkg/croot

CGO_LDFLAGS=-lcroot

CGOFILES=\
	croot_cintex.go \
	croot_genreflex.go \
	croot.go \
	croot_reflex.go


include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
