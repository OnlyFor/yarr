VERSION=2.4
GITHASH=$(shell git rev-parse --short=8 HEAD)

GO_TAGS    = sqlite_foreign_keys sqlite_json
GO_LDFLAGS = -s -w -X 'main.Version=$(VERSION)' -X 'main.GitHash=$(GITHASH)'

export GOARCH      ?= amd64
export CGO_ENABLED  = 1

build_default:
	mkdir -p _output
	go build -tags "$(GO_TAGS)" -ldflags="$(GO_LDFLAGS)" -o _output/yarr ./cmd/yarr

build_macos:
	mkdir -p _output/macos
	GOOS=darwin go build -tags "$(GO_TAGS) macos" -ldflags="$(GO_LDFLAGS)" -o _output/macos/yarr ./cmd/yarr
	cp src/platform/icon.png _output/macos/icon.png
	go run ./cmd/package_macos -outdir _output/macos -version "$(VERSION)"

build_linux:
	mkdir -p _output/linux
	GOOS=linux go build -tags "$(GO_TAGS) linux" -ldflags="$(GO_LDFLAGS)" -o _output/linux/yarr ./cmd/yarr

build_windows:
	mkdir -p _output/windows
	go run ./cmd/generate_versioninfo -version "$(VERSION)" -outfile src/platform/versioninfo.rc
	windres -i src/platform/versioninfo.rc -O coff -o src/platform/versioninfo.syso
	GOOS=windows go build -tags "$(GO_TAGS) windows" -ldflags="$(GO_LDFLAGS) -H windowsgui" -o _output/windows/yarr.exe ./cmd/yarr

serve:
	go run -tags "$(GO_TAGS)" ./cmd/yarr -db local.db

test:
	go test -tags "$(GO_TAGS)" ./...
