pro:
	docker rmi -f web-service:1.0
	docker-compose up
dev:
	go run main.go