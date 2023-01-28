main:
	docker compose run --rm go go run main.go
build:
	docker compose run --rm go go build
download:
	docker compose run --rm go go get
tidy:
	docker compose run --rm go go mod tidy