.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build -o httpapi cmd/httpapi/main.go

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	./httpapi

.PHONY: dub
dub:
	docker compose -f build/docker-compose.yaml up --build

.PHONY: du
du:
	docker compose -f build/docker-compose.yaml up

.PHONY: dd
dd:
	docker compose -f build/docker-compose.yaml down
