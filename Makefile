build:
	./build.sh

docker-build: build
	docker-compose up --build

docker-up:
	docker-compose up

docker-rm:
	docker rm mini_queue_1

default: build