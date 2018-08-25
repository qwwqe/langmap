.PHONY: run
run:
	go run \
		--ldflags "-X github.com/qwwqe/langmap.Version=$$(git describe --tags --always)" \
		cmd/langmap/main.go \
			-f extras/config.json
