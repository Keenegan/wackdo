APP_NAME := wackdo

.PHONY: run
run:
	go run src/main.go

.PHONY: test
test:
	go test ./...

.PHONY: check
check:
	staticcheck ./...

.PHONY: up
up:
	docker-compose up -d --build	

.PHONY: down
down:
	docker-compose down

.PHONY: restart
restart: down up
