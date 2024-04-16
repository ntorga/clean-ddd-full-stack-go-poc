SHELL=/bin/bash
UI_DIR=src/presentation/ui

dev:
	air serve

refresh:
	tmux new-session -d -s hot-reload-go "sleep 1; xdotool search --onlyvisible --name 'chrome|firefox|opera' windowactivate --sync key F5"

build:
	templ generate -path $(UI_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app ./main.go

run:
	./bin/app serve