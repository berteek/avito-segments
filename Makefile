.PHONY: build
build:
	go build -o app cmd/main.go

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
