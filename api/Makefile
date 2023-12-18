include /${PWD}/.env
start:
		@go run .
dev:
		@gin --appPort 8080 --port 8070  --immediate run .
build:
		@go build .
run: build
		./create-transaction