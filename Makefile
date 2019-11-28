.PHONY: build
build:
	@mkdir bin >/dev/null 2>&1 || true
	go build -o bin/helm-take-ownership -ldflags "-X main.date=$$(date "+%Y-%m-%d")"

.PHONY: release
release:
	goreleaser --rm-dist
