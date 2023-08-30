.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build -o cmd/httpapi/httpapi cmd/httpapi/main.go

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	cmd/httpapi/httpapi

.PHONY: dub
dub:
	docker compose -f build/docker-compose.yaml up --build

.PHONY: du
du:
	docker compose -f build/docker-compose.yaml up

.PHONY: dd
dd:
	docker compose -f build/docker-compose.yaml down
