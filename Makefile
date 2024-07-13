.PHONY: all
all:
	go build -ldflags='-s -w' -o ./bin/geomap ./cmd/geomap/main.go

.PHONY: lint
lint:
	goimports-reviser -format ./...
	golangci-lint run

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf ./bin/*
