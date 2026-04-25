.PHONY: run run-cli run-gui build-cli build-gui

run: run-gui

run-cli:
	go run ./cmd/linux/cli/main.go

run-gui:
	go run ./cmd/linux/gui/main.go

