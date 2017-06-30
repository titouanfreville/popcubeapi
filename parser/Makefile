# Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# blame: jnml, labs.nic.cz

default: run

all: simpleparser.go tokenizer.go
	gofmt -l -s -w *.go
	go build $^

run: all
	cd .. && go run test_parser.go

tokenizer.go: simpleparser.l
	golex -o $@ $<

simpleparser.go: simpleparser.y
	goyacc -o $@ $<

clean:
	go clean
	rm -f tokenizer.go simpleparser.go lex.yy.go y.go y.output *~