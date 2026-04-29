.PHONY: run run-cli run-gui

run-cli:
	go run ./cmd/linux/cli/main.go

run-gui:
	go run ./cmd/linux/gui/main.go

run: run-cli