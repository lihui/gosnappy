include $(GOROOT)/src/Make.inc

TARG=lihui/gosnappy

CGOFILES=\
	gosnappy.go

CGO_LDFLAGS=-lsnappy

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w=true *.go

