build:
	./build.sh

docker-build: build
	docker-compose up --build

docker-up: build
	docker-compose up

default: build