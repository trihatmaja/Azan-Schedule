.PHONY: cli

VERSION=`git log --pretty=format:'%H' -n 1`
BUILD=`date +%FT%T%z`

cli:
	go build -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}" -o azan_cli app/cli/main.go
api:
	go build -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}" -o azan_api app/api/main.go
