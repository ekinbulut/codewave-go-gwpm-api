.PHONY: build run clean

build:
	go build -o build/app cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf build

docker.app: docker.app.build
	docker run --rm -d \
		--name dev-rabbitmq-app \
		--network dev-network \
		-p 8080:8080 \
		dev-rabbitmq-app

docker.rabbitmq:
	docker run --rm -d \
		--name dev-rabbitmq \
		--hostname dev-rabbitmq \
		--network dev-network \
		-v ${HOME}/dev-rabbitmq:/var/lib/rabbitmq \
		-v ${PWD}/configs/definitions.json:/opt/definitions.json:ro \
		-v ${PWD}/configs/rabbitmq.config:/etc/rabbitmq/rabbitmq.config:ro \
		-p 5672:5672 \
		-p 15672:15672 \
		rabbitmq:3-management

docker.app.build:
	docker build -t dev-rabbitmq-app .

docker.stop:
	docker stop dev-rabbit docker.app
	
docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network