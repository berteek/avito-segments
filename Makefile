.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build -o app cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	./app

.PHONY: dub
dub:
	docker compose up --build

.PHONY: du
du:
	docker compose up

.PHONY: dd
dd:
	docker compose down
