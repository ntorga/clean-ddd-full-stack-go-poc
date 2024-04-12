dev:
	browser-sync start --port 3000 --proxy localhost:8080 --no-ui &
	air serve
	pkill -f "browser-sync"

build:
	templ generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app ./main.go

run:
	./bin/app serve