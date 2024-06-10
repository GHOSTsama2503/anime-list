install: install-api install-web

install-api:
	cd api && \
	make install

install-web:
	cd web && \
	pnpm run install


run: run-api run-web

run-api:
	cd api && \
	make run

run-web:
	cd web && \
	pnpm run dev


build: build-api build-web

build-api:
	cd api && \
	make build

build-web:
	cd web && \
	pnpm run build


clean-build: clean-build-api clean-build-web

clean-build-api:
	rm -r build/api/ && \
	build-api
