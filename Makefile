.PHONY: all test release clean

GOBIN = build/bin

all:
	go build -trimpath -o $(GOBIN)/stratum-ping

test: all
	go test -count=1 -v ./...

release:
	build/release.sh

clean:
	rm -fr $(GOBIN) build/dist
