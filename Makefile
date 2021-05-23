init:
	docker-compose -f docker/docker-compose.yaml pull

start:
	docker-compose -f docker/docker-compose.yaml up -d nats

stop:
	docker-compose -f docker/docker-compose.yaml stop
	docker-compose -f docker/docker-compose.yaml rm

build:
	go build main.go

listen:
	docker-compose -f docker/docker-compose.yaml up receive send
