build:
	CGO_ENABLED=0 GOOS=linux go build -o build/main main.go

run:
	./build/main

load-test:
	k6 run k6/generate-order-id.js