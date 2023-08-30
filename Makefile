dc:
	docker-compose up --remove-orphans --build

test:
	@go test -v ./...
