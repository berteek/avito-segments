.PHONY: build
build:
	go build -o app cmd/main.go

.PHONY: run
run:
	./app
