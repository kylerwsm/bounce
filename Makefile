.PHONY: build format run

.DEFAULT_GOAL := run

run: format
	go run main.go

format:
	go mod tidy && go mod vendor

push: format
	docker build -t kylerwsm/bounce .
	docker push kylerwsm/bounce
	echo 'y' | docker system prune -a
