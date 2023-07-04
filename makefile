build:
	go build -o build/app cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf build

.PHONY: build run clean

docker:
	docker build -t app .
	docker run -p 8080:8080 -d app