.PHONY: all test clean

GOBIN = build/bin

all: 
	build/env.sh go build -v -o build/bin/stratum-ping

debug: 
	build/env.sh go get -race -v ./...

test: all
	build/env.sh go test -count=1 -timeout 0 -v ./...

clean:
	build/env.sh go clean -cache -modcache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*
