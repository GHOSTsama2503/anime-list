install:
	go mod download && go mod verify

run:
	go run cmd/main.go

build:
	go build -ldflags "-s -w" -o ../build/api/anime-list cmd/main.go
