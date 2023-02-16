.PHONY: 
		build
		run

build:
	docker-compose build film-app

run:
	docker-compose up film-app

test:
	go test -v ./...

swag:
	swag init -g cmd/main.go