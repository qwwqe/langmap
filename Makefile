VERSION = $(shell git describe --tags --always)

CONFIG = extras/config.json
INPUT = cmd/langmap/main.go
LDFLAGS = -X github.com/qwwqe/langmap.Version=$(VERSION)

langmap:
	go build -o $@ -ldflags "$(LDFLAGS)" $(INPUT)

.PHONY: run
run:
	go run -ldflags "$(LDFLAGS)" $(INPUT) -config $(CONFIG) $(LANGMAP_ARGS)

.PHONY: clean
clean:
	rm -rf langmap
