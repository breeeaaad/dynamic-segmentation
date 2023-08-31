build:
	docker-compose up --build

run: build
	docker-compose up --remove-orphans

stop: 
	docker-compose down --remove-orphans