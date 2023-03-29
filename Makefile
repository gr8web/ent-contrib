#!/usr/bin/make -f

ENTSRC := $(wildcard ent/schema/*.go)

ent/ent.go: GOOS=
ent/ent.go: GOARCH=
ent/ent.go:
	$(RM) -r $(filter-out ent/schema ent/generate.go,$(wildcard ent/*))
	go generate ./...

clean:
	$(RM) -r $(filter-out ent/schema ent/generate.go,$(wildcard ent/*))

fmt:
	gofmt -s -w $(SRC)

