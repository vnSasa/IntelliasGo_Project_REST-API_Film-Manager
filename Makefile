.PHONY: 
		build
		run

build:
	docker-compose build film-app

run:
	docker-compose up film-app