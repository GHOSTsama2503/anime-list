install: install-api install-web

install-api:
	go mod download && go mod verify

install-web:
	cd web && \
	pnpm run install


run: run-api run-web

run-api:
	go run cmd/main.go

run-web:
	cd web && \
	pnpm run dev


build: build-api build-web

build-api:
	go build -ldflags "-s -w" -o build/api/anime-list cmd/main.go

build-web:
	cd web && \
	pnpm run build


clean-build: clean-build-api clean-build-web

clean-build-api:
	rm -r build/api/ && \
	build-api
