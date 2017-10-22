.PHONY: cli

cli:
	go build -o azan_cli app/cli/main.go
api:
	go build -o azan_api app/api/main.go
