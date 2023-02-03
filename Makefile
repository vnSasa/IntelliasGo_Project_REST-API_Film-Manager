.PHONY: 
		build
		run
		migrate

build:
	docker-compose build film-app

run:
	docker-compose up film-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:110513@localhost:5436/postgres?sslmode=disable' up